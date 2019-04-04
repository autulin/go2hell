package bencher

import (
	"fmt"
	"sync"
	"time"
)

type Bench struct {
	Qps        int // 每秒请求数
	Concurrent int // 并发度
	Duration   int // 持续时间 单位秒
	JobImpl    IJob
}

type IJob interface {
	Init()
	Exe()
	End()
}

func NewBench(qps, concurrent, duration int, JobImpl IJob) *Bench {
	return &Bench{
		Qps:        qps,
		Concurrent: concurrent,
		Duration:   duration,
		JobImpl:    JobImpl,
	}
}

func (b *Bench) Start() {
	fmt.Printf("====> init job \n")
	b.JobImpl.Init()

	fmt.Printf("====> start job with qps: %v, concurrent: %v, duration: %v s\n", b.Qps, b.Concurrent, b.Duration)
	intervalNanoSec := time.Second.Nanoseconds() / int64(b.Qps) / int64(b.Concurrent)
	ticker := time.NewTicker(time.Duration(intervalNanoSec))

	countSum := int64(b.Duration) * time.Second.Nanoseconds() / intervalNanoSec
	count := int64(0)

	var wg sync.WaitGroup
	for {
		<-ticker.C

		for i := 0; i < int(b.Concurrent); i++ {
			wg.Add(1)
			go func() {
				b.JobImpl.Exe()
				wg.Done()
			}()
		}

		count++
		if count >= countSum {
			break
		}
	}

	fmt.Printf("====> wait to end... \n")
	wg.Wait()
	ticker.Stop()
	b.JobImpl.End()
	fmt.Printf("====> job done! \n")
}
