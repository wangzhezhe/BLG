package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func system(s string) {
	cmd := exec.Command("/bin/sh", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func main() {
	system("who")
	system("pwd")
	system("cd ./src")
	system("pwd")

}
