package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//this could get the current dir path
	current_dir, _ := os.Getwd()
	fmt.Println(current_dir)
	file_path := current_dir + "/temp.txt"

	//reading file into the memory once time
	dat, err := ioutil.ReadFile(file_path)
	check(err)
	//the type of data is []uint
	fmt.Println(dat)
	fmt.Println(string(dat))

	//use this way to obtain the os.File value
	//if you want to do some other file operation
	f, err := os.Open(file_path)
	//if there are some err it will be shown
	check(err)

	b1 := make([]byte, 5)
	//read the first 5 byte from the file to n1
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//read the file from the No.6 position
	//read two byte out from that position
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes after %d position : %s\n", n2, o2, string(b2))

	//the similar function can be implement by the io package
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes after %d : %s\n", n3, o3, string(b3))

	//??
	_, err = f.Seek(0, 0)
	check(err)

	//The bufio package implements a buffered reader that may be useful
	//both for its efficiency with many small reads and because of the additional reading methods it provides.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	//fmt.Println(string(b4))
	fmt.Printf("5 bytes : %s\n", string(b4))

	//close the file when you are done
	f.Close()

}
