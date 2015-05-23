//easy and clear
//channels are pipes between concurent goroutines
//you can send mesg to a channel and recieve it by another channel
package main

import "fmt"

func main() {
	//create a new channel
	messages := make(chan string)

	//send the message into the channel
	//its an anonymous function
	go func() { messages <- "ping" }()
	//msg recieve the info from a channel
	msg := <-messages
	fmt.Println(msg)
	//by default
	//sends and receives block until both the sender and receiver are ready
	//this allow us synchronization between the goroutines

}
