package main

import "fmt"

//closure 一般中文翻译成闭包
//理解闭包最好的方式就是 将闭包函数看成一个类
//intSeq看成一个闭包类 i 是里面的一个全局变量
//每次返回的值会对i进行更新
//貌似闭包也能定义在结构体当中

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}

}

func main() {
	//所谓的函数式的编程
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//这里相当于重新初始化 了一次 就又返回了1
	newInts := intSeq()
	fmt.Println(newInts())

}

/*

补充 函数式编程 (比较透彻)
https://github.com/ZhangzheBJUT/blog/blob/master/closure.md
函数式编程是一种编程模型，他将计算机运算看做是数学中函数的计算，并且避免了状态以及变量的概念。
闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)。这个从字面上很难理解，特别对于一直使用命令式语言进行编程的程序员们。
在函数式编程中 函数是基本单位 是一等公民
其实理解闭包最方便的方法就是将闭包函数看成一个类，一个闭包函数调用就是实例化一个类（在Objective-c中闭包就是用类来实现的）

*/
