package main

import "fmt"

func main() {
	//using this two kind of syntax to create a map
	mb := map[int]string{1: "hello", 2: "go"}
	for key, value := range mb {
		fmt.Println("the key is: ", key, "the value is: ", value)

	}
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	//remove key-value from a map
	delete(m, "k2")
	fmt.Println("map", m)

	_, prs := m["k2"]
	fmt.Println("prs", prs)

}
