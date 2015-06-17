//执行系统命令 并且使用打开的文件作为标准输入和输出
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// 这种方式可以动态输出
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

func main() {
	f, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	cmd := exec.Command("echo", "hello world")
	cmd.Stdout = f
	cmd.Stderr = f
	cmd.Run()
}
