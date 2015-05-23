package main

import (
	"fmt"
	"time"
)

//input name is done  input type is chan bool
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	//if the message did not send into the channel done
	//the <- done and following code will not be excuted
	//so we could use  this mechnisam to realize the sychonization
	<-done
	fmt.Println("following code be excuted")
}

//注意channel的这种无缓存时候的特性常常可以用来对两个goroutines进行同步
//在其中一个goroutine中 写入 done<-true
//之后如果没有读出信息 这里就会一直阻塞
//之后执行完其他部分的代码之后 再将 done中的内容写出来 <-done这样就完成了同步 控制了goroutin的执行点
