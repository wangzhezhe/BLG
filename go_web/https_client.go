package main

import (
	//"io"
	//"log"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//x509.Certificate.
	pool := x509.NewCertPool()
	caCertPath := "certs/ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	//pool.AppendCertsFromPEM(caCrt)
	pool.AddCert(caCertPath)

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	//if use https://10.10.xx.xx
	//the error : panic: Get https://10.10.105.204:2379/v2/keys/foo: x509:
	//cannot validate certificate for 10.10.105.204 because it doesn't contain any IP SANs
	//需要先添加好 etcd 的映射关系
	resp, err := client.Get("https://etcdhost:2379/v2/keys")
	if err != nil {
		panic(err)
	}
	//fmt.Println(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
