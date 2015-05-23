package main

import (
	"fmt"
)

// `` show the raw string
func main() {
	s1 := "hello\nworld!"
	s2 := `hello\nworld!`
	fmt.Println(s1)
	fmt.Println(s2)
}
