package main

import (
	"fmt"
)

var ch3 chan int
var ch4 chan int

var chs = []chan int{ch3, ch4}
var numbers = []int{1, 2, 3, 4, 5}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}

func main() {
	select {
	case getChan(0) <- getNumber(2):
		fmt.Println("1th case is selected")
	case getChan(1) <- getNumber(3):
		fmt.Println("2th case is selected")
	default:
		fmt.Println("default!")
	}

	//每次传入到channel中的值可能都不一样
	chanCap := 5
	ch7 := make(chan int, chanCap)
	for i := 0; i < chanCap; i++ {
		select {
		case ch7 <- 1:
		case ch7 <- 3:
		case ch7 <- 5:
		}
	}

	for i := 0; i < chanCap; i++ {
		fmt.Printf("%v\n", <-ch7)
	}

}
