package main

//using flag package
import (
	"flag"
	"fmt"
	//	"os"
)

func main() {
	//fmt.Println(os.Args)
	//当前设置好的都是默认值 之后直接把对应的参数添加进来即可
	//三个参数分别是 flag 默认值 help信息
	ok := flag.Bool("ok", false, "is ok")
	id := flag.Int("id", 0, "id")
	port := flag.String("port", ":8080", "http listen port")
	var name string
	flag.StringVar(&name, "name", "123", "name")

	flag.Parse()

	fmt.Println("ok:", *ok)
	fmt.Println("id:", *id)
	fmt.Println("port:", *port)
	fmt.Println("name:", name)

}
