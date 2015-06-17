package main

import (
	"fmt"
)

func main() {
	var i interface{} = 123
	fmt.Printf("%T->%d\n", i, i)
	j := i.(int)
	fmt.Printf("%T->%d\n", j, j)

}
