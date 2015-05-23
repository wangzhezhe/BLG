package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {

	//注意 time.millisecond的类型为time.Duration 所以在数字相乘的时候 要先将数字转化成为time.Duration的类型
	fmt.Println(reflect.TypeOf(time.Millisecond))
	//在var中的使用for循环 直接i:=0就会报错？
	var i = 0
	for i = 0; i < 10; i++ {
		x := time.Duration(rand.Intn(100))
		time.Sleep(x * time.Millisecond)

		fmt.Println(x)
	}

}
