package main

import "fmt"

type rect struct {
	width, height int
}

//format for function:
//func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {..}
//format for methods:
//func (r ReceiverType) funcName(parameters) (results)
//attention :A method is a function with an implicit first argument, called a receiver
//method is invoked by .  try to use receiver_instance.methodsname
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("preim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())

}
