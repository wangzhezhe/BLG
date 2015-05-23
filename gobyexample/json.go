// detail info refer to the
//https://gobyexample.com/json
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//We’ll use these two structs to demonstrate encoding and decoding of custom types below.
type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//First we’ll look at encoding basic data types to JSON strings. Here are some examples for atomic values.
	//the basic type to json
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	//此时实际存在[]中的是asc码 marshal
	fmt.Println(fltB)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//And here are some for slices and maps, which encode to JSON arrays and objects as you’d expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	//这个就是常见的json形式了
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//json package could atomatically encode the custome data type
	//The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	//decoding json into the go values
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//fmt.Println(byt)
	//we shuld put the package where json data could put the decode data
	//This map[string]interface{} will hold a map of strings to arbitrary data types
	//We need to provide a variable where the JSON package can put the decoded data. This map[string]interface{} will hold a map of strings to arbitrary data types.

	var dat map[string]interface{}

	//Here’s the actual decoding, and a check for associated errors.
	//this is the actuall decoding
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)
	//In order to use the values in the decoded map, we’ll need to cast them to their appropriate type
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	//We can also decode JSON into custom data types
	//This has the advantages of adding additional type-safety to our programs
	//and eliminating the need for type assertions when accessing the decoded data
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//We can also stream JSON encodings directly to os.Writers
	//like os.Stdout or even HTTP response bodies.
	//不再需要本地结构作为中介
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
	//fmt.Println(enc)
}
