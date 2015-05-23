package main

import "fmt"

//the basic using of slice
func main() {
	//slice have fixed number object
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	//add new element
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	//copy operation
	copy(c, s)
	fmt.Println("cpy:", c)

	//截取slice >=position2 <posiiton5
	l := s[2:5]
	fmt.Println("sl1 ", l)

	// <position 5
	l = s[:5]
	fmt.Println("sl2 ", l)
	//>=position 2
	l = s[2:]
	fmt.Println("sl3 ", l)
	//create an arry
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// create a slice which has 3 element every element is a slice
	// attention the creation of two dimension slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
