package main

import (
	"fmt"
	"sort"
)

//the type defiend by user
//example of custome the sort using go
type ByLength []string

//the Sort function in sort package use the interface which have the
// Len Swap and Less three funciton
// if we want to use Sort we shuld implements this three functions by ourselves

/*

We implement sort.Interface - Len, Less, and Swap -
on our type so we can use the sort package’s generic Sort function.
Len and Swap will usually be similar across types and Less will hold the actual custom sorting logic.
In our case we want to sort in order of increasing string length,
so we use len(s[i]) and len(s[j]) here.

*/
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//如果len str1 < len str2 就被认为是正确的排列
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "bananas", "kiwi"}
	//注意这里还要强制转换一下
	//将fruits类型强制转化为ByLength类型
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

/*

in general

By following this same pattern of creating a custom type,
implementing the three Interface methods on that type,
and then calling sort.Sort on a collection of that custom type,
we can sort Go slices by arbitrary functions


*/
