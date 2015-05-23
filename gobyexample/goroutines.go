//the basic example of goroutines
package main

import "fmt"

func f(from string) {
	for i := 0; i < 20; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	go f("goroutine")
	//anonymous invoking (the last part is the input parameter)
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
