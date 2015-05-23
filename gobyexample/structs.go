package main

import "fmt"

type person struct {
	name string
	age  int
}

//the struct could also include the anonnymous fieds
type student struct {
	person
	specialstring string
}

func main() {
	//show the different way to initial a struct
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	p := &s
	fmt.Println(p)

	sp := s
	fmt.Println(s.name)
	//struct are mtable
	sp.age = 51
	fmt.Println(sp.age)

	//the way to access the anonymous fields in a struct
	stu := student{person{"Amy", 24}, "computing"}
	stu.person.name = "John"
	fmt.Println(stu.person.name)
}
