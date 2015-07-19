package main

import (
	"fmt"
)

type Test []interface{}

func main() {
	test := make(Test, 5)
	test[0] = "a"
	test[1] = 2
	test[2] = true
	test[3] = []byte("test")

	for index, element := range test {
		if value, ok := element.(string); ok {
			fmt.Printf("test[%d] is type of string the value is %v\n", index, value)
		} else if value, ok := element.([]byte); ok {
			fmt.Printf("test[%d] is type of string the []byte is %v\n", index, value)
		}
	}
}
