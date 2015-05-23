package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

//the type of parameter is a pointer
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	//the parameted is a pointer the contends of pointer pointed is 0
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	//get the address of i
	fmt.Println("pointer:", &i)
}
