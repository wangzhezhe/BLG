package BasicUsings

import "fmt"

func MapBasicOperation() {

	//Create and traverse map
	mb := map[int]string{1: "hello", 2: "go"}
	for key, value := range mb {
		fmt.Println("the key is: ", key, "the value is: ", value)
	}

	//Create map using keyword make
	//Attention:if not init the map before the assignment operation
	//there will be the panic:assignment to entry in nil map
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

	newm := make(map[string]string)
	newm["key1"] = "string1"
	fmt.Println(newm)

	//change the value in map
	//Attention:the addressability of the golang map if the value is the struct type
	//refer to https://wangzhezhe.github.io/2017/01/22/golang-map-addressbility/
	newm["key1"] = "string2"
	fmt.Println(newm)

	//check the value in map
	v, ok := newm["key2"]
	if ok {
		fmt.Println("the value with index key2 exists in map", v)
	} else {
		fmt.Println("none value")
	}
}
