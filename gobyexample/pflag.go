package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

var flagvar int

func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
}
func main() {
	fmt.Println("value: ", flagvar)
}
