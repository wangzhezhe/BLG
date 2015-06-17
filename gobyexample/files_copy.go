package main

import (
	"io"
	"os"
)

func main() {
	//注意不要打开两次
	dst, err := os.OpenFile("destfile", os.O_WRONLY|os.O_CREATE, 0644)
	src, err := os.Open("./dat2")
	if err != nil {
		panic(err)
	}
	//  dst, err := os.Open("destfile")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
	src.Close()
	dst.Close()

	return
}
