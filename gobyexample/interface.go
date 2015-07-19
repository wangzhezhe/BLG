package main

//interfaces is based on methods
//interfaces are named collection of method signature
//grouping and naming related sets of methods into interfaces.
import (
	"fmt"
	"math"
)

//the contends in interface is the methods without receiver part
type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

//different implementation of interfaces
func (s square) area() float64 {
	return s.height * s.width
}

func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	//	return 2 * math.pi * c.radius
	return 2
}

//based on function not class
//the variable g has the geometry type
//so we can call methods in the interface in this method
//when invoke this function the input parameter should be the receiver of methods
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circle{radius: 5}

	//attention:The circle and square struct types
	//both implement the geometry interface
	//so we can use instances of these structs as arguments to measure.
	//the input argument is receiver
	measure(s)
	measure(c)
}
