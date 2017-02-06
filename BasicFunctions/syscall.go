package BasicFunctions

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func SystemExeca(s string) {
	cmd := exec.Command("/bin/sh", "-c", s)
	fmt.Println(s)
	out, err := cmd.StdoutPipe()
	go func() {
		o := bufio.NewReader(out)
		for {
			line, _, err := o.ReadLine()
			if err == io.EOF {
				break
			} else {
				fmt.Println(string(line))
			}
		}
	}()
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func SystemExecb(s string) {
	cmd := exec.Command("/bin/sh", "-c", s)
	fmt.Println(s)
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stdout = os.Stdout
	read, _ := cmd.StdoutPipe()
	defer read.Close()

	go func() {
		io.Copy(os.Stdout, read)
	}()

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", out.String())
}

func SystemExecc(s string) {
	cmd := exec.Command("/bin/sh", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}
