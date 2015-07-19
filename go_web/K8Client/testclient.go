package main

import (
	"fmt"
	//"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	//"encoding/json"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client"
	//"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	//"os"
)

//using https way in kube connection
//refer to the https://github.com/GoogleCloudPlatform/kubernetes/issues/10159#issuecomment-113955582
func main() {
	//	ns := api.NamespaceDefault
	//	_, err := os.Open("cert/apiserver.crt")

	//	if err != nil {
	//		panic(err)
	//	}

	config := &client.Config{
		//Host: "http://121.40.171.96:8080",
		Host:    "https://10.10.105.250:8081",
		Version: "v1beta3",
		TLSClientConfig: client.TLSClientConfig{
			// Server requires TLS client certificate authentication
			//CertFile: "cert/server.crt",
			// Server requires TLS client certificate authentication
			//KeyFile: "cert/server.key",
			// Trusted root certificates for server
			CAFile: "cert/ca.crt",
		},
		BearerToken: "abcdTOKEN1234",
	}
	client, err := client.New(config)
	if err != nil {
		// handle error
	}

	//label := map[string]string{}
	//label["name"] = "t2-lgwar-001"
	//result, err := client.Pods("default").List(nil, nil)
	//labels.selectfromset
	//result, err := client.Services("default").List(nil)
	//label["name"] = "test-logtest.war-1.1.1"
	result, err := client.Services("default").List(nil)
	//jstr, _ := json.Marshal(result)
	//fmt.Println(string(jstr))
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}

}
