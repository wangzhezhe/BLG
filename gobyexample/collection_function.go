//collection function的使用还是比较灵活的
//go语言不支持泛型
//函数中传入的参数 可以使某个集合类型
package main

import (
	"fmt"
	"strings"
)

//vs中的第几个值等于t
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	//如果t不被包含在vs中就返回-1
	return -1
}

//t是否包含在vs中
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//注意传过来的第二个参数是一个函数
//如果slice中的任意一个值 满足函数f的要求 就返回true
//将函数作为一个参数来传递
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//只有slice中的元素全部满足函数f的要求 才返回true 否则就返回false
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}

	return true
}

//通过Filter对slice中的元素进行过滤
//先创建一个长度为0的 vsm 之后遇到符合要求的vsm就添加在后面 最后返回vsm
func Filter(vs []string, f func(string) bool) []string {
	vsm := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsm = append(vsm, v)
		}
	}

	return vsm
}

//map中的第一个元素是slice中的位置
//map中的第二个元素是对该位置上的元素进行函数f运算之后得到的函数
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	//传递的参数是匿名函数
	//slice中任何一个string的前缀包含p
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	//slice中每一个string包含字母e 则返回true
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	//注意这里是将函数传递进去 而并非进行真正的调用
	fmt.Println(Map(strs, strings.ToUpper))
}
