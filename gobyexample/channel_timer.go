//这个可以结合timeout一起来使用
//timer(定时器) is also based on channel
//感觉底层像是在channel同步的机制上建立起来的
package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	//时间到了之后（过了2s）就执行
	<-timer1.C
	fmt.Println("Time 1 expired")

	timer2 := time.NewTimer(time.Second * 1)

	go func() {
		//过了1s就执行
		<-timer2.C
		fmt.Println("Time 2 expired")
	}()

	//time.Sleep(time.Second * 3)
	//stop函数的返回值是bool类型
	//it retures the true if the call stops the timer
	//if the timer is already expired or stopped （已经执行完或者已经停止了）
	//it returns false
	stop2 := timer2.Stop()
	fmt.Println(stop2)
	if stop2 {
		fmt.Println("Time 2 stopped")
	}
}
