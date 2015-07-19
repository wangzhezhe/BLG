package main

import (
	"fmt"
	//"time"
)

func main() {
	c1 := make(chan string)
	//time.Sleep(time.Second)
	//c2 := make(chan bool)
	//c1 <- "test default"
	go func() {
		//time.Sleep(time.Second * 1)
		c1 <- "test"
	}()
	//time.Sleep(time.Second * 1)
	//using flag to adjust if it was exit from the for circulation
	//flag := 0

	//if you want to use the synchornization and the select default at the same time
	//maybe you should use a for circulation to do some check

	//在多层循环的时候 可以通过打标签的方式来控制跳转的层次 之后与break结合在一起使用
	//比如这个 在select中的break A的时候 这样直接跳转出来的就是最外层的A循环
	//如果不加上break标签的话 直接跳转出来的就是最里面一层的 select 循环
A:
	for i := 0; ; i++ {
		//time.Sleep(time.Second * 1 / 100)
		select {
		case res := <-c1:
			fmt.Println(res)
			//flag = 1U
			break A
		default:
			fmt.Println("no message received: ", i)

		}
		//if flag == 1 {
		//break
		//}
	}

	/*
		msg := true
		select {
		case c2 <- msg:
			fmt.Println(c2)
		default:
			fmt.Println("error")
		}

		select {
		case msg := <-c1:
			fmt.Println(msg)
		case res := <-c2:
			fmt.Println(re	s)
		default:
			fmt.Println("no activity")
		}
	*/

}
