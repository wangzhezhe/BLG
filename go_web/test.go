package main

import (
	//	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

var ncpu int32

func plus(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("the default value:", ncpu)
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	fmt.Println("the type is :", reflect.TypeOf(jsonStr))
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)
	fmt.Println(os.Getwd())
	fmt.Println([]byte("<br>"))

}
