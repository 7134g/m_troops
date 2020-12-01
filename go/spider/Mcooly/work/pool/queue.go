package pool

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/setting"
	"net/url"
	"sync"
	"sync/atomic"
	"time"
)

const stop = true

type Queue struct {
	// Threads defines the number of consumer threads
	Threads           int
	storage           queue.Storage
	activeThreadCount int32
	threadChans       []chan bool
	lock              *sync.Mutex

	// my add
	queueMaxCount int          // 队列最大值
	SpiderName    string       // 爬虫名
	taskCount     int          // 任务数
	ActiveCount   int          // 活动数
	activeLock    sync.RWMutex // 计算活动的锁
	taskLock      sync.RWMutex // 计算活动的锁
}

// New creates a new queue with a Storage specified in argument
// A standard InMemoryQueueStorage is used if Storage argument is nil.
func New(threads int, s queue.Storage, maxCount int) (*Queue, error) {
	if s == nil {
		s = &queue.InMemoryQueueStorage{MaxSize: maxCount}
	}
	if err := s.Init(); err != nil {
		return nil, err
	}
	q := Queue{
		Threads:     threads,
		storage:     s,
		lock:        &sync.Mutex{},
		threadChans: make([]chan bool, 0, threads),
	}
	q.queueMaxCount = maxCount
	return &q, nil
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	s, _ := q.Size()
	return s == 0
}

// AddURL adds a new URL to the queue
func (q *Queue) AddURL(URL string) error {
	q.lock.Lock()
	q.lock.Unlock()
	u, err := url.Parse(URL)
	if err != nil {
		return err
	}
	r := &colly.Request{
		URL:    u,
		Method: "GET",
	}
	d, err := r.Marshal()
	if err != nil {
		return err
	}
	return q.AddTask(d)
}

// AddRequest adds a new Request to the queue
func (q *Queue) AddRequest(r *colly.Request) error {
	d, err := r.Marshal()
	if err != nil {
		return err
	}
	if err := q.AddTask(d); err != nil {
		return err
	}
	q.lock.Lock()

	//q.taskCount++
	for _, c := range q.threadChans {
		c <- !stop
	}
	q.threadChans = make([]chan bool, 0, q.Threads)
	q.lock.Unlock()
	return nil
}

func (q *Queue) AddTask(d []byte) error {
	var count int
	for {
		count = q.GetTaskCount()
		if count >= q.queueMaxCount {
			logs.Log.Warning(q.SpiderName + " 当前任务超过最大值，开始沉睡等待")
			time.Sleep(setting.SPIDERSLEEPTIME)
		} else {
			//logs.Log.Debug(s.SpiderName + " 队列值小于最大值，继续执行任务...")
			break
		}
	}
	q.addTaskCount()
	return q.storage.AddRequest(d)
}

// Size returns the size of the queue
func (q *Queue) Size() (int, error) {
	return q.storage.QueueSize()
}

// Run starts consumer threads and calls the Collector
// to perform requests. Run blocks while the queue has active requests
func (q *Queue) Run(c *colly.Collector) error {
	wg := &sync.WaitGroup{}
	for i := 0; i < q.Threads; i++ {
		wg.Add(1)
		go func(c *colly.Collector, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				if q.IsEmpty() {
					if q.activeThreadCount == 0 {
						break
					}
					ch := make(chan bool)
					q.lock.Lock()
					q.threadChans = append(q.threadChans, ch)
					q.lock.Unlock()
					action := <-ch
					if action == stop && q.IsEmpty() {
						break
					}
				}
				q.lock.Lock()
				atomic.AddInt32(&q.activeThreadCount, 1)
				q.lock.Unlock()
				rb, err := q.storage.GetRequest()
				if err != nil || rb == nil {
					q.finish()
					continue
				}
				r, err := c.UnmarshalRequest(rb)
				if err != nil || r == nil {
					q.finish()
					continue
				}
				q.reduceTaskCount()
				q.addActiveCount()
				q.waitActiveExcute()
				r.Do()
				q.finish()
			}
		}(c, wg)
	}
	c.Wait()
	q.delGetActiveThread()
	wg.Wait()
	return nil
}

func (q *Queue) finish() {
	q.lock.Lock()
	q.activeThreadCount--
	for _, c := range q.threadChans {
		c <- stop
	}
	q.threadChans = make([]chan bool, 0, q.Threads)
	q.lock.Unlock()

	q.reduceActiveCount()
}

func (q *Queue) GetActiveCount() int {
	q.activeLock.Lock()
	defer q.activeLock.Unlock()

	return q.ActiveCount
}

func (q *Queue) addActiveCount() {
	q.activeLock.Lock()
	defer q.activeLock.Unlock()
	q.ActiveCount += 1
}

func (q *Queue) waitActiveExcute() {
	var count int
	for {
		count = q.GetActiveCount()
		if count > setting.GOROUTINECOUNT {
			logs.Log.Warning(q.SpiderName + " 当前任务超过最大值，开始沉睡等待")
			time.Sleep(setting.SPIDERSLEEPTIME)
		} else {
			//logs.Log.Debug(q.SpiderName + " 队列值小于最大值，继续执行任务...")
			break
		}
	}
}

func (q *Queue) reduceActiveCount() {
	q.activeLock.Lock()
	defer q.activeLock.Unlock()
	q.ActiveCount -= 1
}

func (q *Queue) delGetActiveThread() {
	q.activeLock.Lock()
	defer q.activeLock.Unlock()
	q.ActiveCount = 0
}

func (q *Queue) GetTaskCount() int {
	q.taskLock.Lock()
	defer q.taskLock.Unlock()

	return q.taskCount
}

func (q *Queue) addTaskCount() {
	q.taskLock.Lock()
	defer q.taskLock.Unlock()
	q.taskCount++
	return
}

func (q *Queue) reduceTaskCount() {
	q.taskLock.Lock()
	defer q.taskLock.Unlock()
	q.taskCount--
	return
}
