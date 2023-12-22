package clock

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	"github.com/zeromicro/go-zero/core/utils"
	"time"
)

const (
	day = time.Hour * 24
)

type TickerExecute struct {
	ticker *time.Ticker

	executeTime uint // 执行次数

	interval              time.Duration        // 间隔时间
	intervalStatus        bool                 // 间隔时间执行标志
	intervalCalibrateFunc func() time.Duration // 校准时间,使用场景：启动时候，设置的时间小于当前时间
	startTime             *time.Time           // 开始执行时间
	endTime               *time.Time           // 结束执行时间

	hour, minter, second int // 时钟时间
	day                  int // 增减天数

	retryTime time.Duration      // 默认五分钟后重试
	cancel    context.CancelFunc // 取消定时器
}

func NewTickerExecute() *TickerExecute {
	t := &TickerExecute{}
	t.SetRetryTime(time.Minute * 5)
	t.SetIntervalCalibrateFunc(func() time.Duration {
		return t.interval
	})
	return t
}

func (t *TickerExecute) SetStartTime(hour, minter, second int) {
	now := time.Now()
	target := time.Date(now.Year(), now.Month(), now.Day(), hour, minter, second, 0, now.Location())
	t.startTime = &target
}

func (t *TickerExecute) SetEndTime(hour, minter, second int) {
	now := time.Now()
	target := time.Date(now.Year(), now.Month(), now.Day(), hour, minter, second, 0, now.Location())
	t.endTime = &target
}

func (t *TickerExecute) SetDay(day int) {
	t.day = day
}

// SetRetryTime 发生错误时候重试时间
func (t *TickerExecute) SetRetryTime(duration time.Duration) {
	t.retryTime = duration
}

// SetIntervalTimer 设置间隔时间
func (t *TickerExecute) SetIntervalTimer(value time.Duration) {
	t.intervalStatus = true
	t.interval = value
}

// SetTimerAlarmClock 定时闹钟
func (t *TickerExecute) SetTimerAlarmClock(hour, minter, second int) {
	t.hour = hour
	t.minter = minter
	t.second = second
}

func (t *TickerExecute) SetIntervalCalibrateFunc(f func() time.Duration) {
	t.intervalCalibrateFunc = f
}

func (t *TickerExecute) getIntervalTimer(now time.Time) time.Duration {
	// 判断是否处于运行时间范围
	if t.startTime != nil {
		startTime := time.Date(now.Year(), now.Month(), now.Day(),
			t.startTime.Hour(), t.startTime.Minute(), t.startTime.Second(), 0, now.Location())
		if startTime.After(now) {
			return startTime.Sub(now)
		}
	}

	if t.endTime != nil {
		endTime := time.Date(now.Year(), now.Month(), now.Day(),
			t.endTime.Hour(), t.endTime.Minute(), t.endTime.Second(), 0, now.Location())
		if endTime.After(now) {
			return t.interval
		} else {
			// 超出今天的执行时间
			startTime := time.Date(now.Year(), now.Month(), now.Day(),
				t.startTime.Hour(), t.startTime.Minute(), t.startTime.Second(), 0, now.Location())

			return startTime.AddDate(0, 0, 1).Sub(now)
		}
	}

	if t.executeTime == 0 {
		// 首次执行
		targetTime := time.Date(now.Year(), now.Month(), now.Day(), t.hour, t.minter, t.second, 0, now.Location())
		if targetTime.After(now) {
			return targetTime.Sub(now)
		} else {
			return t.intervalCalibrateFunc()
		}
	}

	return t.interval
}

func (t *TickerExecute) GetSleepTime() time.Duration {
	now := time.Now()
	targetTime := time.Date(now.Year(), now.Month(), now.Day(), t.hour, t.minter, t.second, 0, now.Location())

	if t.intervalStatus {
		// 设置了自定义间隔时间执行
		return t.getIntervalTimer(now)
	}

	// >
	if targetTime.After(now) {
		// 今天执行
		return targetTime.Sub(now)
	} else {
		// 明天执行
		return targetTime.AddDate(0, 0, 1).Sub(now)
	}
}

type jobFunc func(ctx context.Context, targetTime time.Time) error

type jobTask interface {
	DoTask(ctx context.Context, targetTime time.Time) error
}

func (t *TickerExecute) notify() {
	// todo 通知机器人发生了panic
}

func (t *TickerExecute) Cancel() {
	t.cancel()
}

func (t *TickerExecute) run(f jobFunc, ctx context.Context, targetTime time.Time) {
	t.executeTime++
	if err := f(ctx, targetTime); err != nil {
		logc.Errorf(ctx, "DoTask: %v", err)
		// 五分钟后重试
		t.ticker.Reset(t.retryTime)
		return
	}
}

func (t *TickerExecute) Run(j jobTask) {
	RunSafe(func() {
		// 初始化定时器
		interval := t.GetSleepTime()
		t.ticker = time.NewTicker(interval)
	})

	go RunSafe(func() {
		cancelCtx, cancel := context.WithCancel(context.Background())
		t.cancel = cancel

		for {
			timer := utils.NewElapsedTimer()
			select {
			case <-t.ticker.C:
				ctx := context.Background()
				targetTime := time.Now().AddDate(0, 0, t.day)
				t.run(j.DoTask, ctx, targetTime)
			case <-cancelCtx.Done():
				logc.Infow(cancelCtx, fmt.Sprintf("TickerExecute cancel %T", j), logx.Field("timeout", timex.ReprOfDuration(timer.Duration())))
				return
			}
			logc.Infow(cancelCtx, fmt.Sprintf("TickerExecute Done: %T", j), logx.Field("timeout", timex.ReprOfDuration(timer.Duration())))
			interval := t.GetSleepTime()
			t.ticker.Reset(interval)
		}
	}, t.notify)
}

// RunLastDay 统计昨天
func (t *TickerExecute) RunLastDay(j jobTask) {
	t.day = -1
	t.Run(j)
}
