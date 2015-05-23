//range could provide a iteration of data structure
//we can also use range to receive mutiple values from a channel
//if we need the index we could use this syntax: for i,num :=range nums {}
package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

	nums := []int{2, 3, 4}
	sum := 0
	//the first position is index
	//if you do not use it you could use a '_' to replace it
	for _, sum := range nums {
		sum += sum
	}

	fmt.Println("sum:", sum)
	//index 要从0开始
	for i, num := range nums {
		if num == 2 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("%s -> %s\n", k, v)
	}

	//在这种情况下 会遍历 “test” 串中的每一个字符 index会从0开始
	for i, c := range "test" {
		fmt.Println(i, c)
	}

}
