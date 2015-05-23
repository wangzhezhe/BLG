package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//shows how to dump a string into the file
	d1 := []byte("hello go\n")
	err := ioutil.WriteFile("dat1", d1, 0644)
	check(err)

	//for more granular writes open a file for writing
	f, err := os.Create("dat2")
	fmt.Println(reflect.TypeOf(f))

	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	//Issue a Sync to flush writes to stable storage.
	f.Sync()

	//bufio provides buffered writers
	//in addition to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Println("wrote %d bytes\n", n4)

	//Use Flush to ensure all buffered operations have been applied to the underlying writer.
	w.Flush()
}
