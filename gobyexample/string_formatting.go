//主要是 Printf功能的演示
//这个基本上与C中的用法是一样的 还有一些go本身新加的功能 比如 %v

package main

import (
	"fmt"
	//"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	//%v print go value
	fmt.Printf("%v\n", p)
	//If the value is a struct,
	//the %+v variant will include the struct’s field names.
	fmt.Printf("%+v\n", p)
	//The %#v variant prints a Go syntax representation of the value,
	//i.e. the source code snippet that would produce that value.
	fmt.Printf("%#v\n", p)

	//print the type of value
	//it is more convinient than using reflect
	fmt.Printf("%T\n", p)

}
