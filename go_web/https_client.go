package main

import (
	//"io"
	//"log"
	"crypto/tls"
	"crypto/x509"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"strings"
)

func main() {
	//x509.Certificate.
	//pool := x509.NewCertPool()
	//caCertPath := "etcdcerts/ca.crt"
	//caCertPath := "certs/cert_server/ca.crt"

	//	caCrt, err := ioutil.ReadFile(caCertPath)
	//	if err != nil {
	//		fmt.Println("ReadFile err:", err)
	//		return
	//	}
	//pool.AppendCertsFromPEM(caCrt)
	//pool.AddCert(caCrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool,
			InsecureSkipVerify: true},
		DisableCompression: true,
	}
	url = ""

	reqest, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}
	response, _ := client.Do(reqest)
	reqest.Header.Set("Authorization", "qwertyuiopasdfghjklzxcvbnm1234567890")
	//client := &http.Client{Transport: tr}
	//resp, err := client.Get("https://server:8081")
	//if use https://10.10.xx.xx
	//the error : panic: Get https://10.10.105.204:2379/v2/keys/foo: x509:
	//cannot validate certificate for 10.10.105.204 because it doesn't contain any IP SANs
	//需要先添加好 etcd 的映射关系
	//resp, err := client.Get("https://etcdhost:2379/v2/keys/foo")
	//resp, err := client.Get("https://k8master:8081/api/v1beta3/namespaces/default/pods")
	//resp, err := client.Get("https://Apitransfer:8080/api/v1beta3/namespaces/default/pods")
	//	var clusterinfo = map[string]string{}
	//	clusterinfo["username"] = "yyy"
	//	clusterinfo["password"] = "123456"
	//	clusterinfo["masterip"] = "10.10.105.250"
	//	clusterinfo["cacrt"] = "ca string"
	//	body, _ := json.Marshal(clusterinfo)
	//resp, err := client.Post("https://Apitransfer:10443/v1/application/checkuser", "application/json", strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}
	//fmt.Println(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.Status)
}
