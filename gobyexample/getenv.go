package main

import (
	"fmt"
	"os"
)

//有的时候特别是在容器的环境下
//需要从环境变量中将所有的配置信息读入
//注意良好的设计性 所有的配置信息最好都放在一个配置文件中
func main() {
	env := os.Getenv("HOME")
	fmt.Println(env)
}
