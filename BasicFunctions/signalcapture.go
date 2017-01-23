package BasicFunctions

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func Systemexec(s string) {
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
