package main

import (
	"fmt"
	"reflect"
)

func main() {
	var b byte = 'D'
	fmt.Println(reflect.TypeOf(b))
}
