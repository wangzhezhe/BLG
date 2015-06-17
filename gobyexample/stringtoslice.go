package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//example 1
	var slice []int32 = make([]int32, 5, 10)
	//p是指向slice类型的一个指针
	//这里是模拟一个slice类型 然后转化过来
	p := (*struct {
		array uintptr
		len   int
		cap   int
	})(unsafe.Pointer(&slice))

	//para v will print the go type
	//If the value is a struct,
	//the %+v variant will include the struct’s field names
	fmt.Printf("output:%+v\n", p)

	//example 2
	//声明数组的时候 没有显式指定数组的长度 而是通过特殊的标记 ...
	//告诉编译器 让编译器去自动地计算
	var array = [...]int32{1, 2, 3, 4, 5}
	//数组进行截取 转化为slice 此时自动给len 和 cap 赋了值
	var sliceb = array[2:4]
	fmt.Printf("before modify: array=%+v, slice = %+v\n", array, sliceb)
	//此时 slice 中用的底层数组和 array 是同一个
	//此时 slice[0]的位置 指向的实际是底层数组中 元素3所在的位置
	sliceb[0] = 678
	fmt.Printf("after modify: array=%+v, slice = %+v\n", array, sliceb)

	//example 3
	//注意 如果使用的是append方式生成新的slice
	//就不会有类似的效果 因为采用append方式会分配新的底层数组
	array = [...]int32{1, 2, 3, 4, 5}
	var slicec = array[1:3]
	slicec = append(slicec, 6, 7, 8)

	//可以对比这次结果 发现 两次array的输出是一样的 并没有因为 slice的修改而对array造成影响
	fmt.Printf("before modify: array=%+v, slice = %+v\n", array, sliceb)
	slicec[0] = 123
	fmt.Printf("after modify: array=%+v, slice = %+v\n", array, sliceb)

	//注意 从slice 生成slice 的时候原理也是类似的 直接用slicea=sliceb[para1:para2]的语法
	//使用的是同一个底层的数组 要是通过append方式 则会重新分配底层数组

}

/* 输出：
output:&{array:8725963136 len:5 cap:10}
before modify: array=[1 2 3 4 5], slice = [3 4]
after modify: array=[1 2 678 4 5], slice = [678 4]
before modify: array=[1 2 3 4 5], slice = [3 4]
after modify: array=[1 2 3 4 5], slice = [3 4]
*/
