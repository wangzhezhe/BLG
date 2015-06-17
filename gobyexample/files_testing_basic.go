package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {

	dirname := "file_temp3"

	//判断目录是否存在 Stat函数返回一些对文件的基本描述信息 比如名字大小
	info, err := os.Stat(dirname)

	//目录存在就报错
	if err == nil {
		fmt.Println("the dir already exist")
		fmt.Println("size:", info.Size())

	}

	//创建文件夹
	err = os.Mkdir(dirname, 0777)
	if err != nil {
		panic(err)
	}
	//在文件夹中创建文件
	//这个是对 openfile 函数的封装 默认是创建 权限为0666的文件
	f, err := os.Create(dirname + "/" + "dat2")
	fmt.Println(reflect.TypeOf(f))
	if err != nil {
		panic(err)
	}

	//可以从目录上看到文件被创建出来 之后几秒钟后又被删除

	time.Sleep(time.Second * 2)
	//文件包括文件里面的内容都被删除
	err = os.RemoveAll(dirname)
	if err != nil {
		panic(err)
	}
	//	//将新创建的文件删除
	//	err = os.Remove(dirname + "/" + "dat2")
	//	if err != nil {
	//		panic(err)
	//	}
	//	time.Sleep(time.Second * 2)
	//	//将文件夹删除
	//	err = os.Remove(dirname)
	//	if err != nil {
	//		panic(err)
	//	}

}
