package pool

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

/*
基本流程如下：
初始化一个协程池(pool)，调用pool.Submit(task func())提交Task到pool，判断pool是否有空闲的worker，如果有则取-个work来处理Task，
如果没有，判断pool是否满了，是，则阻塞等待pool中有worker时再处理Task;否，则新建一个worker来处理Task，处理完成后回收worker到pool。
调用者可通过defer pool.Close()，关闭pool, 清理过期worker的goroutine退出; 调用pool.ReSize(size);重置pool中worker数量，调用后立即生效。
如果未调用pool.Close(), 将保证pool至少有一个worker可用, 以便下一次任务调用;
*/

const (
	// 默认协程清理时间
	DefaultCleanIntervalTime = 2 * time.Second

	//工作状态
	Waiting int = 0
	Working int = 1
	Idling  int = 2
)

// 需要执行的任务
type Task struct {
	TaskFunc func([]interface{})
	Param    []interface{}
}

// 工人
type worker struct {
	task          chan *Task //工人工作任务，同一时间只能做一个
	startWaitTime time.Time  //工人开始等待的时间
	status        int        //工作状态
	mtx           sync.RWMutex
}

// 设置工作状态
func (w *worker) changeStatus(status int) {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	w.status = status
}

// 获取当前工作状态
func (w *worker) currentStatus() int {
	w.mtx.RLock()
	defer w.mtx.RUnlock()
	return w.status
}

// 厂房
type Pool struct {
	capacity       int32         //容量
	pruneDuration  time.Duration //清理时间间隔
	expiryDuration time.Duration //worker过期时间
	workers        []*worker     //worker集合
	release        chan struct{} //关闭池信号
	lock           sync.Locker   //同步操作锁
	once           sync.Once     //单例
	closed         bool          //协程池是否已关闭
	wg             sync.WaitGroup
	mtx            sync.Mutex
	Ctx            context.Context // 控制是否终止协程
}

// 正在工作的工人数
func (p *Pool) Running() int32 {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	return int32(len(p.workers))
}

// 厂房能容纳同时工作的工人数量
func (p *Pool) Cap() int32 {
	return atomic.LoadInt32(&p.capacity)
}

// 新建一个厂房
func NewPool(size int32, prune bool, expiry_time time.Duration) (*Pool, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	p := &Pool{
		capacity:       size,                   //厂房容量
		release:        make(chan struct{}, 1), //监察停止监控信号
		pruneDuration:  2 * time.Millisecond,   //清理间隔时间
		expiryDuration: expiry_time,            //工人空闲下班时间
		lock:           newSpinLock(),
		Ctx:            ctx, //控制是否停止发包
	}
	//启动定期清理空闲工人的协程,如果协程池本身比较小，不建议开启清理
	if prune && size > 200 {
		go p.periodicallyPurge()
	}
	return p, cancel
}

// 定期清理过期的工人, 在厂房关闭前至少保留一个工人在工作
func (p *Pool) periodicallyPurge() {
	heartbeat := time.NewTicker(p.pruneDuration)
	defer heartbeat.Stop()
	var (
		currentTime time.Time
		unInitTime  time.Time
	)
	for !p.closed {
		select {
		case <-heartbeat.C:
			if p.Running() <= 1 {
				continue
			}
			var pruned = 0
			for {
				pruned = 0
				currentTime = time.Now()
				p.lock.Lock()
				for i, w := range p.workers {
					if w.currentStatus() == Waiting {
						//新建的Worker时间是未初始化的
						if w.startWaitTime == unInitTime {
							continue
						} else if currentTime.Sub(w.startWaitTime) >= p.expiryDuration {
							close(w.task)
							w.changeStatus(Idling)
							p.workers = append(p.workers[:i], p.workers[i+1:]...)
							pruned++
							break
						}
					}
				}
				p.lock.Unlock()
				if p.Running() <= 1 || pruned == 0 {
					break
				}
			}
		case <-p.release:
			return
		}
	}
}

// 提交一个任务
func (p *Pool) Submit(task *Task) error {
	//没有厂房或厂房已关闭，那么报错
	if p == nil {
		return errors.New("this pool is nil")
	}
	if p.closed {
		return errors.New("this pool has been closed")
	}
	//安排一个工人，把任务交给它
	w := p.getWorker()
	w.changeStatus(Working)
	p.wg.Add(1)
	w.task <- task
	return nil
}

// 安排一个工人
func (p *Pool) getWorker() *worker {
	var wk *worker
	if p.Running() > p.Cap() { //如果超过上限，则清理掉一部分已完成任务的worker
		for !p.closed {
			p.lock.Lock()
			for i, w := range p.workers {
				//如果是等待任务的状态，可以清理
				if w.currentStatus() == Waiting {
					close(w.task) //从队列中移除前必须先close chan
					p.workers = append(p.workers[:i], p.workers[i+1:]...)
					break
				}
			}
			p.lock.Unlock()
			if p.Running() <= p.Cap() {
				break
			}
		}
	}
	if p.Running() == p.Cap() { //如果已到容量上限，则寻找空闲的worker
	loop:
		for !p.closed {
			p.lock.Lock()
			for _, w := range p.workers {
				//如果是等待任务的状态，可以分配
				if w.currentStatus() == Waiting {
					wk = w
					p.lock.Unlock()
					break loop
				}
				//等待100微秒
				time.Sleep(time.Duration(100) * time.Microsecond)
			}
			p.lock.Unlock()
		}
	} else { //如果没到容量上限，则添加新的workers
		//招募一个工人
		wk = &worker{
			task:   make(chan *Task, 1),
			status: Waiting,
			mtx:    sync.RWMutex{},
		}
		p.lock.Lock()
		p.workers = append(p.workers, wk)
		p.lock.Unlock()
		//让工人开始接收任务
		p.startWork(wk)
	}
	return wk
}

// 让工人开始接收任务
func (p *Pool) startWork(worker *worker) {
	go func() {
		for {
			select {
			case f, ok := <-worker.task:
				if !ok {
					return
				}
				f.TaskFunc(f.Param)
				//标记工作完成
				p.wg.Done()
				//设置为等待状态
				worker.changeStatus(Waiting)
				//开始等待计时
				worker.startWaitTime = time.Now()
			case <-p.Ctx.Done():
				//工作取消，设置为等待状态
				worker.changeStatus(Waiting)
			}
		}
	}()
}

// 调整厂房工人的数量
func (p *Pool) ReSize(size int32) {
	//数量刚刚好
	if size == p.Cap() {
		return
	}
	//保证至少有一个工人在工作
	if size <= 0 {
		size = 1
	}
	atomic.StoreInt32(&p.capacity, size)
}

// 等待正在工作中的工人完成任务
func (p *Pool) Wait() {
	p.wg.Wait()
}

// 关闭协程池，让所有worker退出运行以防goroutine泄露,必须先调用Wait
func (p *Pool) Close() {
	p.once.Do(func() {
		//通知清理协程退出
		close(p.release)
		//将所有工人的任务清空，退出doWork的for range循环
		p.lock.Lock()
		for _, worker := range p.workers {
			close(worker.task)
		}
		p.lock.Unlock()
		p.workers = p.workers[:0]
		p.closed = true
	})
}
