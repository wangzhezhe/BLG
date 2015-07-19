package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	sign := make(chan byte, 2)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i

			time.Sleep(1 * time.Second)
		}

		close(ch)
		fmt.Println("The channel is closed.")
		sign <- 0
	}()

	go func() {
		//这个循环会一直尝试从ch中读取信息出来 即使ch已经被发送端关闭
		//但还是可以读信息出来 最后当ok 为false的时候 说明已经没有数据从ch中读出
		//跳出循环 注意这种判断方式
		for {
			fmt.Printf("before extract channel len: %v ,", len(ch))
			e, ok := <-ch
			fmt.Printf("channel value: %d if extract ok :(%v) after extraction channel len : %v channel cap : %v \n", e, ok, len(ch), cap(ch))
			if !ok {
				break
			}

			time.Sleep(2 * time.Second)
		}
		fmt.Println("Done.")
		sign <- 1
	}()
	//要是不添加两次取值的操作的话 主进程就会马上结束 这里相当于是实现了一个
	//同步的操作 等待两个go func都结束之后 再结束主进程 注意这种技巧
	<-sign
	<-sign
}

/*output:
before extract channel len: 1 ,channel value: 0 if extract ok :(true) after extraction channel len : 0 channel cap : 1
before extract channel len: 1 ,channel value: 1 if extract ok :(true) after extraction channel len : 0 channel cap : 1
before extract channel len: 1 ,channel value: 2 if extract ok :(true) after extraction channel len : 0 channel cap : 1
before extract channel len: 1 ,channel value: 3 if extract ok :(true) after extraction channel len : 0 channel cap : 1
The channel is closed.
before extract channel len: 1 ,channel value: 4 if extract ok :(true) after extraction channel len : 0 channel cap : 1
before extract channel len: 0 ,channel value: 0 if extract ok :(false) after extraction channel len : 0 channel cap : 1
Done.
*/
