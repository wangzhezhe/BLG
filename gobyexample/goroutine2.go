package main

import (
	"fmt"
	//"runtime"
	"math/rand"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		//runtime.Gosched()
		x := time.Duration(rand.Intn(100))
		time.Sleep(x * time.Millisecond)
		fmt.Println(i, " ", s)
	}
}

func main() {

	go say("a")
	go say("b")
	go say("c")
	say("world")

}
