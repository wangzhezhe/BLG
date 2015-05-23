/*
	url 的基本格式
	scheme://host[:port#]/path/.../[?query-string][#anchor]
	scheme         指定低层使用的协议(例如：http, https, ftp)
	host           HTTP服务器的IP地址或者域名
	port#          HTTP服务器的默认端口是80，这种情况下端口号可以省略。如果使用了别的端口，必须指明，例如 http://www.cnblogs.com:8080/
	path           访问资源的路径
	query-string   发送给http服务器的数据 注意这个
	anchor         锚


*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	u, _ := url.Parse("10.180.87.216:9090/testparameter")
	//u, _ := url.Parse("http://www.baidu.com")
	q := u.Query()
	q.Set("username", "user")
	q.Set("password", "passwd")
	u.RawQuery = q.Encode()

	response, err := http.Get(u.String())
	if err != nil {
		fmt.Println("error occured")

	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

}
