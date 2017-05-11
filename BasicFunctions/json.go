//Transfer json info into golang struct:
//https://mholt.github.io/json-to-go/

//refer to the simplejson to encoding the json info
//using simplejson when you need to unmarshal only part of the data
//http://stackoverflow.com/questions/21432848/go-json-with-simplejson
//problem in \x00 error/??

//the initial character should be the capital form
//http://stackoverflow.com/questions/25187718/invalid-character-x00-after-top-level-value
//more info http://attilaolah.eu/2013/11/29/json-decoding-in-go/

package BasicFunctions

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

//using the struct tags to match the json data with the struct
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func JsonBasicOperation() {
	//First we’ll look at encoding basic data types to JSON strings. Here are some examples for atomic values.
	//the basic type to json
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)

	fmt.Println(fltB)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	//And here are some examples for slices and maps, which encode to JSON arrays and objects as you’d expect.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

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

	//Decoding json into the go values
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
	str := `{"page": 1, "fruits": ["apple", "peach"],"abc":"der"}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	//We can also stream JSON encodings directly (without using the binary data) to os.Writers
	//like os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
	//fmt.Println(enc)

	//other example is to using the simple json package to marchal only the small part of all json data
	//http://stackoverflow.com/questions/21432848/go-json-with-simplejson
}

type Inner struct {
	Innera string `json:"innera"`
	Innerb string `json:"innerb"`
	Innerc string `json:"innerc"`
}

type OuterA struct {
	Outera string `json:"outera"`
	Outerb string `json:"outerb"`
	Outerc string `json:"outerc"`
	Inner
}

type OuterB struct {
	Outera        string `json:"outera"`
	Outerb        string `json:"outerb"`
	Outerc        string `json:"outerc"`
	Innerinstance Inner  `json:"inner"`
}

type OuterC struct {
	Outera        string `json:"outera"`
	Outerb        string `json:"outerb"`
	Outerc        string `json:"outerc"`
	Innerinstance Inner
}

//more info here(about the valid tag type) attention to the essence of the tag
//http://stackoverflow.com/questions/10858787/what-are-the-uses-for-tags-in-go
func EmbedJsonTransfermation() {
	fmt.Println("show the struct embedying in json transfermation")

	inner := Inner{
		Innera: "ina",
		Innerb: "inb",
		Innerc: "inc"}

	outerA := OuterA{
		Outera: "oua",
		Outerb: "oub",
		Outerc: "ouc",
		Inner:  inner}

	outerB := OuterB{
		Outera:        "oua",
		Outerb:        "oub",
		Outerc:        "ouc",
		Innerinstance: inner}

	outerC := OuterC{
		Outera:        "oua",
		Outerb:        "oub",
		Outerc:        "ouc",
		Innerinstance: inner}

	outerStringa, err := json.Marshal(outerA)
	if err != nil {
		fmt.Printf("json error %+v", err)
	}
	fmt.Println(string(outerStringa))

	outerStringb, err := json.Marshal(outerB)
	if err != nil {
		fmt.Printf("json error %+v", err)
	}
	fmt.Println(string(outerStringb))

	outerStringc, err := json.Marshal(outerC)
	if err != nil {
		fmt.Printf("json error %+v", err)
	}
	fmt.Println(string(outerStringc))

	/*
		output:
		   {"outera":"oua","outerb":"oub","outerc":"ouc","innera":"ina","innerb":"inb","innerc":"inc"}
		   {"outera":"oua","outerb":"oub","outerc":"ouc","inner":{"innera":"ina","innerb":"inb","innerc":"inc"}}
		   {"outera":"oua","outerb":"oub","outerc":"ouc","Innerinstance":{"innera":"ina","innerb":"inb","innerc":"inc"}}

	*/

}
