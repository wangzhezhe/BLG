//很好的学习tar功能的例子
//http://www.nljb.net/default/Golang-%E5%8E%8B%E7%BC%A9-%E8%A7%A3%E5%8E%8B-Tar.Gz/

//这个例子的主要功能是把 testfile目录中的 内容打包好之后 放到tar文件夹下面 所生成的压缩文件中
package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	//file write
	//首先要有tar目录 目录不会被自动创建 但是文件可以被自动创建
	//创建出最后要打包成的tar文件
	fw, err := os.Create("tar/lin_golang_src.tar.gz")
	//type of fw is *os.File
	fmt.Println(reflect.TypeOf(fw))
	if err != nil {
		panic(err)

	}
	defer fw.Close()

	//gzip writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	//tar write
	//这里最后生成的是.tar.gz类型的文件
	//.tar.gz又添加了一层压缩 而 .tar仅仅是打包 没有添加压缩
	tw := tar.NewWriter(gw)
	defer tw.Close()

	//打开文件夹
	dir, err := os.Open("testfile/")
	if err != nil {
		panic(nil)
	}
	defer dir.Close()

	//读取文件列表
	//如果 n>0 返回最多n个 n<=0 以slice的形式返回全部的的信息
	//fis的类型为[]os.FileInfo
	fis, err := dir.Readdir(0)
	//fis的类型为 []os.FileInfo
	fmt.Println(reflect.TypeOf(fis))
	if err != nil {
		panic(err)
	}

	//遍历文件列表 每一个文件到要写入一个新的*tar.Header
	for _, fi := range fis {
		//如果有文件夹的时候就递归 暂时省略
		//判断
		if fi.IsDir() {
			continue
		}

		//打印文件名称
		fmt.Println(fi.Name())

		//打开文件 open当中是 目录名称/文件名称 构成的组合
		fr, err := os.Open(dir.Name() + "/" + fi.Name())
		if err != nil {
			panic(err)
		}
		defer fr.Close()

		//信息头部 生成tar文件的时候要先写入tar结构体
		//内建函数 new(T) 分配了0值 填充了T类型的存储空间 并且返回其地址
		//返回的地址是一个*T类型的值 即使说返回了一个指针
		h := new(tar.Header)
		fmt.Println(reflect.TypeOf(h))
		h.Name = fi.Name()
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()

		//将信息头部的内容写入
		err = tw.WriteHeader(h)
		if err != nil {
			panic(err)
		}

		//将每个打开文件的 内容写入tar文件 注意tar writer对gzip writer做了一层包装
		//最前面仅仅是生成了一个 tar writer 在循环中 循环到每个文件都要生成一个 file writer 将内容写入到tar中
		//copy(dst Writer,src Reader)
		_, err = io.Copy(tw, fr)
		if err != nil {
			panic(err)
		}

	}

	fmt.Println("tar.gz OK")
}
