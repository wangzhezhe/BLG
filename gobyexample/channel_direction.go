package main

import "fmt"

//the channel pings only recieve the messages
func ping(pings chan<- string, msg string) {
	//this channel only receive mesage
	pings <- msg
}

//the first channel is pings which only send messages out（<-on the left）
//pings后面的三个部分 <- chan string 全是用于定义的语句
//表示说 pings 是一个只能用于向外发送信息的channel
//the second channel pongs only receive message in (<-on the right)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed messages")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
