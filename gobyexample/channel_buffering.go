package main

import "fmt"

func main() {
	//we make a channel of strings buffering up to 2 values
	//put mutiple values into the channels
	//if you want to get the value from the channel immediately
	//maybe you should use buffering
	messages := make(chan string, 2)
	//the buffered in channel is FIFO
	messages <- "buffered"
	messages <- "chanenl"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	/*
	    the following formate without buffering will deadlock
	   	c1 := make(chan string)
	   	c1 <- "test"
	   	select {
	   	case res := <-c1:
	   		fmt.Println(res)
	   	}

	*/

	//this kind of way is ok even if the buffer number is 1
	c1 := make(chan string, 1)
	c1 <- "test"
	select {
	case res := <-c1:
		fmt.Println(res)
	}

	//attention:
	//by default the channel is unbuffered
	//meaning that they will only accept sends if there is a corresponding receive ready to receive the sent value
}
