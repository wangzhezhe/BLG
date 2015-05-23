//in the previous example we use the mutex to synchonized the process
//we could also use the built-in goroutings and channels to achieve the same result

//questions here??

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ops int64 = 0
	//两个channel所接受的类型
	//分别是 readOp和writeOp 两个类型的引用
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		//put the state only in one goroutines
		// this goroutines repeatedly select on the reads and write channels
		// and send value to the response channel
		var state = make(map[int]int)

		//注意这里相当于是while是不断进行着循环的
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
				//fmt.Println(read.resp)
			case write := <-writes:
				state[write.key] = write.val
				//fmt.Println(write.val)
				write.resp <- true
			}
		}
	}()

	//start 100 goroutines  send to the reads channel(接受的是结构体类型)
	//receiving the results from the provided resp channel
	for r := 0; r < 100; r++ {
		go func() {
			//注意这里也是不断在进行着循环
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}

				//只有read.resp返回了结果之后 才能再往下走
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	//start 10 writes using simillar ways
	for w := 0; w < 10; w++ {
		go func() {
			//通过新启动的一个goroutines来不断地进行循环
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	//if do not stop here the goroutines will not be excuted
	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops: ", opsFinal)
}
