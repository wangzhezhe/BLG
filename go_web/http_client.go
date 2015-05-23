package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	//url := "http://10.10.103.86:8080/api/v1beta3/namespaces/default/pods"
	//url := "http://10.10.103.86:8080/api/v1beta3/namespaces/default/replicationcontrollers"
	//url := "http://10.10.103.86:8080/api/v1beta3/namespaces/default/services"
	url := "http://10.211.55.5:8080/v1/testbuild"
	body, err := ioutil.ReadFile("./Dockerfileupload/dockerfile1")
	if err != nil {
		panic(err)
	}

	//reqest, _ := http.NewRequest("GET", url, nil)
	reqest, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}

	//reqest.Header.Set("Content-Type", "application/json")

	response, _ := client.Do(reqest)

	//body为 []byte类型
	returnbody, _ := ioutil.ReadAll(response.Body)
	status := response.StatusCode
	fmt.Println(string(returnbody), status)
}
