package main

import (
	"fmt"
	"log"
	//	"reflect"
	//	"syscall"
)

func main() {
	//log与fmt.println混用 输出的顺序可能会混乱
	log.Printf(" %s", "1.1.1.1002")
	log.Printf(" %s", "aaa")
	//	sigTerm := syscall.Signal(15)
	fmt.Println("HAHAHAHAHAHAHAHAHAHA")
	//log.Println(reflect.TypeOf(sigTerm))
	fmt.Println("ahahahaha")

}
