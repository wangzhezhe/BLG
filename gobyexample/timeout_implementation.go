//利用go的channel来实现限时的操作 很简洁
//主要是基于select
//如果先收到资源所发来的信息 就正常
//如果到了限定的时间还没有收到资源所发来的请求
//时间信息的那个case就会传过来 之后对应的执行就结束
//比如访问某个资源 超过了多长多长 时间  就返回错误 怎样 怎样。。。
//注意time.After函数的应用

package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	//start a new goroutines
	go func() {
		//do sth
		time.Sleep(time.Second * 2)
		//after excution send info to c1
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
		//要是超出了时间 就发送信息给 time.After 这样就相当于超时的操作了
		//time after返回的本来就是一个channel的类型
	case <-time.After(time.Millisecond * 1):
		fmt.Println("time out")
	}

	c2 := make(chan string, 1)
	go func() {

		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("time out")
	}
}
