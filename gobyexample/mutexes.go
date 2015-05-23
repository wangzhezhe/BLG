//For more complex state we can use a mutex to safely access data across multiple goroutines.
//通常的同步是使用goroutine来进行 但也可以用atomic机制来进行 具体见（https://gobyexample.com/atomic-counters）
//使用atomic(sync/atomic package)操作只能处理些比较简单的同步的问题 复杂的情况需要通过 mutexes 来进行
//其实有点类似通常的那种 读写 加锁 同步的问题

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	//the mutex will synchonized the state
	var mutex = &sync.Mutex{}

	var ops int64 = 0

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			//every time we chose a key randomly
			//and access to it
			//the mutex could ensure the exclusive access to the state
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				//将cpu让出 防止starve
				runtime.Gosched()
			}

		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			//for each read process
			//we use the similar way to ensure the exclusive access
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()

			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
