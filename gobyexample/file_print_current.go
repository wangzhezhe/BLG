package main

//import (
//	"os"
//	"os/exec"
//	"path/filepath"
//)

//func main() {
//	file, _ := exec.LookPath(os.Args[0])
//	path, _ := filepath.Abs(file)
//	println(path)
//}

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func dir() {
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
}

//为何显示出来的是一个奇怪的目录？？
///var/folders/fw/8qn9ghz50xsd7h7b7sjjf3wr0000gn/T/go-build334712953/command-line-arguments/_obj/exe
func main() {
	dir()
}
