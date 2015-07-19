package main

import (
	"fmt"
	"time"
)

//方式一 记录两个时间的时间差
//方式二 使用timer
func main() {
	k := time.Now()
	fmt.Println(k)
	//	d, _ := time.Duration("5m")

	//	fmt.Println(k.Add(d))

}
