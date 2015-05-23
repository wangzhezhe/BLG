package main

import (
	"fmt"
	"log"
	"net/http"
	//	"strings"
	"crypto/md5"
	"html/template"
	"io"
	"os"
	"strconv"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	//解析参数 默认情况不会对参数进行解析
	r.ParseForm()
	//服务器端打印
	fmt.Println(r.Form)
	//得到链接后面的各种信息
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("host", r.URL.Host)
	fmt.Println("path", r.URL.Path)
	//return a map
	fmt.Println("Query:", r.URL.Query())
	fmt.Println(r.Form["url_long"])

	//	for k, v := range r.Form {
	//		fmt.Println("key:", k)
	//		fmt.Println("val:", strings.Join(v, ""))
	//	}
	//写入到w 是输出到客户端 可以向writter中写入指定的信息
	fmt.Fprintf(w, "hello from server!")
	//w.Write("nnn")
	s := r.Body
	fmt.Println("the body of request :", s)
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	//设置访问的路由 对路径处理函数进行注册 sayhelloName是具体进行处理的函数
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/upload", upload)
	//设置监听的端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
