//Rate limiting is an important mechanism for controlling resource utilization
//and maintaining quality of service
//速率限制 用来限定发送流量和接受流量的速率

package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		request <- i
	}

	//don't receive the data anymore
	close(request)

	//receive the data per 200 ms
	limiter := time.Tick(time.Millisecond * 1000)

	for req := range request {
		//every 200ms the limiter start oncetime
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	//前面三个都是在极短的时间里写入进来的 之后就是每隔200ms写入一次
	//激活burstLimiter 这个channel的方式有两个 一个是前几个的burst形式的写入
	//另一个就是每隔200ms的激活一次
	//就是所谓了burst request
	for i := 0; i < 2; i++ {
		burstyLimiter <- time.Now()
	}
	//Every 200 milliseconds we’ll try to add a new value to burstyLimiter,
	//up to its limit of 3

	//没隔200ms会有一个突发的启动
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	//simulate the request info
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	//For the result
	//the first five is outputed every 1 second
	//the next three is outputed immediately couse of the burstable rate limiting
	//the next second are served with ~200ms delays each.

}
