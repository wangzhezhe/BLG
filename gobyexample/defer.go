package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/defer.txt")
	/*
		the defer part will be executed at the end of the enclosing function (main),
		after writeFile has finished.

	*/

	defer closeFile(f)
	writeFile(f)
	fmt.Println(f.Name())
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "test data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
