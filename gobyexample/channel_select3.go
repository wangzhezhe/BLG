package main

import (
	"fmt"
	"time"
)

func main() {
	sign := make(chan int)
	chtemp := make(chan int, 5)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 300)
			chtemp <- i
		}
		close(chtemp)

	}()
	var e int
	ok := true

	//new 一个新的channel返回 注意这里要提前声明好
	t := time.After(time.Second)
	go func() {
		for {
			select {
			case <-t:
				fmt.Println("time out")
				ok = false
				break
				//注意这里是使用 = 而不是 :=

			default:
				e, ok = <-chtemp
				fmt.Printf("value : %v \n", e)
				if !ok {
					break
				}
			}
			if !ok {
				sign <- 1
				break
			}
		}
	}()
	<-sign
}
