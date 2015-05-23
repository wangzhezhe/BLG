package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)

func listcontainers(url string, port int) {

	client := &http.Client{}
	str_port := strconv.Itoa(port)
	u_list := url + ":" + str_port + "/containers/json?all=1"

	fmt.Println(u_list)
	reqest, err := http.NewRequest("GET", u_list, nil)

	if err != nil {
		panic(err.Error())
	}

	response, _ := client.Do(reqest)

	body, _ := ioutil.ReadAll(response.Body)

	status := response.StatusCode
	if status == 200 {
		fmt.Println("show container info")
		fmt.Println(string(body))

	} else {
		fmt.Println("error, the status code is ", status)
	}

}

func createcontainer(url string, port int, filename string) string {
	client := &http.Client{}
	str_port := strconv.Itoa(port)
	u_list := url + ":" + str_port + "/containers/create"

	fmt.Println(u_list)

	//read info from temp.json
	byt, err := ioutil.ReadFile(filename)

	//	fmt.Println("the type is :", reflect.TypeOf(byt))
	//	fmt.Println(byt)
	//  js, err := simplejson.NewJson(byt)
	//	fmt.Println(js)
	if err != nil {
		panic(err.Error())
	}

	//the type of byte should be []uint8
	reqest, _ := http.NewRequest("POST", u_list, bytes.NewBuffer(byt))
	reqest.Header.Set("Content-Type", "application/json")

	response, _ := client.Do(reqest)

	//the type of body is []uint8
	//high overhead
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(reflect.TypeOf(body))
	//body_decoder := json.Decoder(body)

	//fmt.Println(string(body))
	//dec := json.NewDecoder(bytes.NewReader(response.Body))

	status := response.StatusCode

	fmt.Println(status)

	if status == 201 {
		fmt.Println("the container is created successfully")
	} else {
		fmt.Println("error in creating the error code is ", status)
	}

	var info_map map[string]string

	//using json.decode
	if err := json.Unmarshal(body, &info_map); err != nil {
		panic(err)
	}

	fmt.Println("the id is :", info_map["Id"])

	return string(info_map["Id"])

}

func startcontainer(url string, port int, id string) {
	client := &http.Client{}
	str_port := strconv.Itoa(port)
	u_list := url + ":" + str_port + "/containers/" + id + "/start"

	reqest, _ := http.NewRequest("POST", u_list, nil)
	reqest.Header.Set("Content-Type", "application/json")

	response, _ := client.Do(reqest)

	//the type of body is []uint8
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	status := response.StatusCode
	if status == 204 {
		fmt.Println("the container is started successfully")
	} else {
		fmt.Println("error in creating the error code is ", status)
	}

}

func deletecontainer(url string, port int, id string) {
	client := &http.Client{}
	str_port := strconv.Itoa(port)
	d_container := url + ":" + str_port + "/containers/" + id

	reqest, _ := http.NewRequest("DELETE", d_container, nil)

	fmt.Println(d_container)
	d_response, _ := client.Do(reqest)

	body, _ := ioutil.ReadAll(d_response.Body)

	status := d_response.StatusCode

	if status == 204 {
		fmt.Println("the container is deleted succesfully")
		fmt.Println(string(body))

	} else {
		fmt.Println("error, the status code is ", status)
	}

}

func main() {

	host_str := "http://10.211.55.5"
	host_port := 2375

	//relative path:http://godoc.org/github.com/kardianos/osext
	//json_create := "/Users/Hessen/goworkspace/go_web/temp.json"
	//get current dir
	current_dir, _ := os.Getwd()
	json_create := current_dir + "/temp.json"
	fmt.Println("current dir: ", current_dir)
	listcontainers(host_str, host_port)
	id := createcontainer(host_str, host_port, json_create)
	startcontainer(host_str, host_port, id)
	time.Sleep(time.Second * 2)
	//deletecontainer(host_str, host_port, id)
}
