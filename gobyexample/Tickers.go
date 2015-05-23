//timers are for when you want to do sth once in the future
//tickers are for when you want to sth repeatedly at regular intervals
//attention:
//Tickers use a similar mechanism to timers: a channel that is sent values

package main

import (
	"fmt"
	"time"
)

func main() {
	//注意时间单位 time.Second 以及 time.Millisecond
	//每隔500ms进行一次迭代
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	//ticker could also be stopped just like the timer
	//once the ticker is stopped it will not
	//receive any values on its channel
	//if the main thread is finished the ticker goroutines
	//are also expired?

	//ticker.Stop()
	//time.Sleep(time.Millisecond * 1500)

	fmt.Println("Ticker stopped")

}
