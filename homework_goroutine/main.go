package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
		1.开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
		2.开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
		3.主 goroutine 从resultChan取出结果并打印到终端输出
*/

type job struct {
	num int64
}

type result struct {
	job *job
	sum int64
}

func creatNum(x chan<- *job) {
	defer wg.Done()
	for {
		rand.Seed(time.Now().UnixNano())
		num := rand.Int63()
		newJob := &job{
			num: num,
		}
		x <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func toSum(y <-chan *job, rs chan<- *result) {
	defer wg.Done()
	for {
		jobtmp := <-y
		n := jobtmp.num
		var sum = int64(0)
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		var newResult = &result{
			sum: sum,
			job: jobtmp,
		}
		rs <- newResult
	}
}

var wg sync.WaitGroup
var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func main() {
	wg.Add(1)
	go creatNum(jobChan)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go toSum(jobChan, resultChan)
	}

	for ret := range resultChan {
		fmt.Printf("生成的随机数：%d, 和：%d\n", ret.job.num, ret.sum)
	}
}
