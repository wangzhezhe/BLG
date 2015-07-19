package main

import (
	//"io"
	//"log"
	"crypto/tls"
	//	"crypto/x509"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	//"strings"
)

func login(user, pw, masterip, cacertLoc string) bool {

	caCertPath := cacertLoc // load ca file
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return false
	}

	// use map as struct
	var clusterinfo = url.Values{}
	//var clusterinfo = map[string]string{}
	clusterinfo.Add("userName", user)
	clusterinfo.Add("password", pw)
	clusterinfo.Add("masterIp", masterip)
	clusterinfo.Add("cacrt", string(caCrt))

	data := clusterinfo.Encode()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true},
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}

	url := "https://10.10.105.135:8443/user/checkAndUpdate"
	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Set("Authorization", "qwertyuiopasdfghjklzxcvbnm1234567890")

	resp, err := client.Do(reqest)

	if err != nil {
		//panic(err)
		fmt.Println(err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return resp.StatusCode == 200
}

func main() {
	userName := "yyy"
	pw := "123456"
	masterip := "121.40.171.96"
	certLoc := "certs/ca.crt"
	if !login(userName, pw, masterip, certLoc) {
		fmt.Println("invalid username or password")
		os.Exit(1)
	}
	fmt.Println("valid user and password, continue")

}
