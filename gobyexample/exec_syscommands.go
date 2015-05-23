//执行系统命令 并且使用打开的文件作为标准输入和输出
package main

import (
	"fmt"
	"os"
	"os/exec"
)

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
