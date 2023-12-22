package clock

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
	"time"
)

type stu struct {
}

func (s *stu) DoTask(ctx context.Context, targetTime time.Time) error {
	fmt.Println("tasking......", targetTime)
	time.Sleep(time.Second * 1)
	fmt.Println("task done......", targetTime)

	return nil
}

func TestTickerExecute_Run(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	s := &stu{}

	t.Run("3 second", func(t *testing.T) {
		ticker := NewTickerExecute()
		ticker.SetIntervalTimer(time.Second * 3)
		ticker.Run(s)
	})

	t.Run("每小时的30分，60分执行", func(t *testing.T) {
		ticker := NewTickerExecute()
		ticker.SetIntervalTimer(time.Second * 3)
		ticker.intervalCalibrateFunc = func() time.Duration {
			var interval time.Duration
			now := time.Now()
			if now.Minute()/30 == 1 {
				interval = time.Duration(60-now.Minute())*time.Minute - time.Duration(now.Second())*time.Second
			} else {
				interval = time.Duration(30-now.Minute())*time.Minute - time.Duration(now.Second())*time.Second
			}

			return interval
		}
		ticker.Run(s)
	})

	t.Run("8点到22点执行", func(t *testing.T) {
		ticker := NewTickerExecute()
		ticker.SetIntervalTimer(time.Second * 3)
		ticker.SetStartTime(16, 45, 0)
		ticker.SetEndTime(16, 54, 0)
		ticker.Run(s)
	})

	select {
	case <-ctx.Done():
		fmt.Println("stop")
		cancel()
		return
	}

}

func TestTickerExecute_RunLastDay(t *testing.T) {
	logx.Disable()
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	s := &stu{}

	t.Run("last day", func(t *testing.T) {
		ticker := NewTickerExecute()
		ticker.RunLastDay(s)

		count := ticker.executeTime
		go func() {
			for {
				if count != ticker.executeTime {
					fmt.Println(ticker.day)
					ticker.day++
					count = ticker.executeTime
				}
				time.Sleep(time.Second * 3)
				fmt.Println("sleep 3 s")
				ticker.ticker.Reset(time.Second * 3)

			}
		}()

	})

	select {
	case <-ctx.Done():
		fmt.Println("stop")
		cancel()
		return
	}
}
