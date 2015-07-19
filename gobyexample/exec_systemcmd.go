package main

import (
	//	"bytes"
	//	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	//	s "strings"
)

func sys(s string) {
	cmd := exec.Command("/bin/sh", "-c", s)
	fmt.Println(s)
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stdout = os.Stdout
	read, _ := cmd.StdoutPipe()
	defer read.Close()
	//r := bufio.NewReader(read)
	//line, _, err := r.ReadLine()
	//fmt.Println(line)
	go func() {
		io.Copy(os.Stdout, read)
	}()

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", out.String())
}

func main() {
	sys("echo hello")

}
