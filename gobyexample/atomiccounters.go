//在go中通常使用channel来进行同步
//当然也可以有其他的方式 比如采用sync/atomic包来进行

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0
	//用于模拟增加数字的操作 每一个gorotines可能会隔1ms左右启动一次
	for i := 0; i < 50; i++ {
		go func() {
			for {
				//这个表示采用原子化的方式执行增量的操作
				atomic.AddUint64(&ops, 1)
				//这个表示允许其他进程来执行
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)
}
