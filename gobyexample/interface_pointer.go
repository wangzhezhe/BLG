package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type SortStringA [3]string
type SortStringB [3]string

func (self SortStringA) Len() int {
	return len(self)
}

func (self SortStringA) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self SortStringA) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}
func (self SortStringA) Sort() {
	sort.Sort(self)
}

func (self *SortStringB) Len() int {
	return len(self)
}

func (self *SortStringB) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self *SortStringB) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self *SortStringB) Sort() {
	//调用sort包中的 Sort方法 传入的参数只要是一个实现了 sort.interface的类型的实例即可
	sort.Sort(self)
}

func main() {

	sa := SortStringA{"2", "3", "1"}
	sb1 := &SortStringB{"2", "3", "1"}
	sb2 := SortStringB{"2", "3", "1"}

	fmt.Println(reflect.TypeOf(sa))
	sorta, ok := interface{}(sa).(Sortable)
	fmt.Println(reflect.TypeOf(sorta))

	fmt.Println(ok) //output:true
	sa.Sort()
	fmt.Printf("SortStringA after sort %v:\n", sa)
	sort.Sort(sa)
	fmt.Printf("SortStringA after sort %v:\n", sa)

	//在golang 的源码包中
	fmt.Println(reflect.TypeOf(sb1))
	sort.Sort(sb1)
	sorted := sort.IsSorted(sb1)
	fmt.Printf("sb1 (type:SortStringB) after sort %v :, is sorted : %v \n", *sb1, sorted)

	sb2.Sort()
	fmt.Printf("sb2 (type:SortStringB) after sort : %v\n", sb2)

}
