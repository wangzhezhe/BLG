//Closing a channel indicates that
//no more values will be sent on it.

package main

import "fmt"

func main() {
	jobs := make(chan string, 5)
	done := make(chan bool)

	//it's another goroutine
	//distinct the extra goroutine with main
	//this goroutine is prepared to recieve the info
	go func() {
		for {
			//the first parameter is the output value
			//the second parameter is the true or false(check if the channel is closed)
			//x,ok<-ch 这个通常用来表示：读取端 是否关闭 如果ok 为true 意味着chanenl尚未被关闭 同时可以读取数据
			//如果ok为false 意味着 channel已经被关闭 不可以再写入数据
			j, more := <-jobs
			if more {
				fmt.Println("received jobs", j)
			} else {
				//if you don't add the close(jobs)
				//this will not be excuted
				fmt.Println(more, j)
				fmt.Println(j)
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
		//<-done
	}()

	for j := 1; j <= 3; j++ {
		jobs <- "test"
		fmt.Println("sent jobs", j)
	}
	//没有数据再发送了就会关闭channel
	close(jobs)
	//if you want to send the info to the channel anymore it will got the error
	//jobs <- 6
	fmt.Println("sent all jobs")
	//the data is stored in channel
	//output the data from the channel
	<-done
}
