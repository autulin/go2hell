package main

import (
	"fmt"
	"github.com/go2hell/bencher"
	"github.com/go2hell/jobs"
	"os"
	"strconv"
	"time"
)

var waitEndCh = make(chan int)

const (
	HelloJob int = 1
)

func main() {

	var clientNum, qps, concurrency, duration, jobType int

	if len(os.Args) < 5 {
		panic("no enough parameter")
	}
	args := os.Args[1:]
	clientNum64, err := strconv.ParseInt(args[0], 10, 0)
	if err != nil {
		println("wrong clientNum. exit")
		return
	} else {
		clientNum = int(clientNum64)
	}
	qps64, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		println("wrong qps. exit")
		return
	} else {
		qps = int(qps64)
	}
	concurrency64, err := strconv.ParseInt(args[2], 10, 0)
	if err != nil {
		println("wrong concurrency. exit")
		return
	} else {
		concurrency = int(concurrency64)
	}
	duration64, err := strconv.ParseInt(args[3], 10, 0)
	if err != nil {
		println("wrong duration. exit")
		return
	} else {
		duration = int(duration64)
	}

	jobType64, err := strconv.ParseInt(args[4], 10, 0)
	if err != nil {
		println("wrong job type. exit")
		return
	} else {
		jobType = int(jobType64)
	}

	st := time.Now().UnixNano()
	for i := 0; i < clientNum; i++ {
		go newTask(i, qps, concurrency, duration, jobType)
	}

	for i := 0; i < clientNum; i++ {
		<-waitEndCh
	}

	fmt.Printf("all finished cost %d ms \n", (time.Now().UnixNano()-st)/time.Millisecond.Nanoseconds())
}

func newTask(clientIdx, qps, concurrency, duration, jobType int) {
	var job bencher.IJob
	switch jobType {
	case HelloJob:
		job = &jobs.HelloJob{}
	default:
		panic("wrong job type")
	}

	bench := bencher.NewBench(qps, concurrency, duration, job)
	bench.Start()

	waitEndCh <- clientIdx
}
