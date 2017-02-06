//using case:
//call on some service , if timeout, do some operation
package BasicFunctions

import (
	"fmt"
	"math/rand"
	"time"
)

func TimeoutOperation() {

	c1 := make(chan string, 1)

	//start a new goroutines
	//after two second, send message to c1
	//the type of x is time.Duration
	x := time.Duration(rand.Intn(5))
	go func() {
		//do sth
		time.Sleep(time.Second * x)
		//after excution send info to c1
		c1 <- "result 1"
	}()
	fmt.Printf("time out after %d second\n", x)
	//if overtime , send message to c1 , print result 1
	//if overtime , send message to c2 , without result 1
	select {
	case res := <-c1:
		fmt.Println(res)

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
