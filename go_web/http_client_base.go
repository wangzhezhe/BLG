//question ???将stdout重定向为response信息？？？
package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	//"reflect"
	"os"
	//	"strings"
)

func main() {

	client := &http.Client{}
	//url := "http://10.10.103.86:8080/api/v1beta3/namespaces/default/pods"
	//url := "http://10.10.103.86:8080/api/v1beta3/namespaces/default/replicationcontrollers"
	//url := "http://10.10.103.86:8080/api/v1beta3/namespaces/default/services"
	//url := "http://10.211.55.8:8080/v1/testbuild"
	//url := "http://10.211.55.5:2375/build?dockerfile=temp_test2/Dockerfile"
	//url := "http://index.docker.io/v1/user"
	url := "http://10.211.55.10:8082/v1/cdf/"
	//	url := "http://10.10.105.45:2376/images/testnewimage/tag"
	//body, err := ioutil.ReadFile("./Dockerfileupload/dockerfile1")
	//if err != nil {
	//	panic(err)
	//}

	reqest, err := http.NewRequest("GET", url, nil)
	//reqest, err := http.NewRequest("POST", url, strings.NewReader("repo=localhost:5000"))

	//reqest, err := http.NewRequest("POST", url, strings.NewReader(string(body)))

	//response, err := http.Post("http://10.10.105.45:2376/images/testnewimage/tag",
	//"application/x-www-form-urlencoded",
	//strings.NewReader("repo=localhost:5000/testnewimage&force=1"))
	if err != nil {
		panic(err)
	}

	//reqest.Header.Set("Content-Type", "application/json")
	//reqest.Header.Set("Content-Type", "application/tar")
	response, _ := client.Do(reqest)

	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)

	//body为 []byte类型
	//returnbody, _ := ioutil.ReadAll(response.Body)
	status := response.StatusCode
	//returnbody 为[]unit8类型
	//fmt.Println(reflect.TypeOf(returnbody))
	fmt.Println(status)
	//fmt.Println(string(returnbody), status)
}
