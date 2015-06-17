//show the client to send the docker build command to build a new image by tar file
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// AuthConfiguration represents authentication options to use in the PushImage
// method. It represents the authentication in the Docker index server.
type AuthConfiguration struct {
	Username      string `json:"username,omitempty"`
	Password      string `json:"password,omitempty"`
	Email         string `json:"email,omitempty"`
	ServerAddress string `json:"serveraddress,omitempty"`
}

type BuildImageOptions struct {
	Name                string            `qs:"t"`
	Dockerfile          string            `qs:"dockerfile"`
	NoCache             bool              `qs:"nocache"`
	SuppressOutput      bool              `qs:"q"`
	RmTmpContainer      bool              `qs:"rm"`
	ForceRmTmpContainer bool              `qs:"forcerm"`
	InputStream         io.Reader         `qs:"-"`
	OutputStream        io.Writer         `qs:"-"`
	RawJSONStream       bool              `qs:"-"`
	Remote              string            `qs:"remote"`
	Auth                AuthConfiguration `qs:"-"` // for older docker X-Registry-Auth header
	//AuthConfigs         AuthConfigurations `qs:"-"` // for newer docker X-Registry-Config header
	ContextDir string `qs:"-"`
}

func SourceTar(filename string) *os.File {
	//"tardir/deployments.tar.gz"
	fw, _ := os.Open(filename)
	//fmt.Println(reflect.TypeOf(fw))
	return fw

}

func main() {

	//dockerhub的认证信息
	auth := AuthConfiguration{
	//	Username:      "wangzhe",
	//	Password:      "3.1415",
	//	Email:         "w_hessen@126.com",
	//	ServerAddress: "https://10.211.55.5",
	}

	//tarStream := SourceTar("./Dockerfileupload/deployments.tar.gz")

	//	opts := BuildImageOptions{

	//		Name:         "testimage",
	//		InputStream:  tarStream,
	//		OutputStream: os.Stdout,
	//		Auth:         auth,
	//		//attention !!! add the full path of Dockerfile after uncompressing
	//		Dockerfile: "/Dockerfile",
	//	}

	client := &http.Client{}
	//url := "http://10.10.103.86:8080/api/v1beta3/namespaces/default/pods"
	//url := "http://10.10.103.86:8080/api/v1beta3/namespaces/default/replicationcontrollers"
	//url := "http://10.10.103.86:8080/api/v1beta3/namespaces/default/services"
	//url := "http://10.211.55.5:8080/v1/testbuild"
	// parameter q could reduce the verbose info
	url := "http://10.211.55.5:2375/build?dockerfile=temp_test2/Dockerfile&q=true&t=easybuild"
	body, err := ioutil.ReadFile("./Dockerfileupload/deployments.tar.gz")
	if err != nil {
		panic(err)
	}

	//reqest, _ := http.NewRequest("GET", url, nil)
	reqest, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}

	//reqest.Header.Set("Content-Type", "application/json")
	reqest.Header.Set("Content-Type", "application/tar")
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(auth)
	reqest.Header.Set("X-Registry-Config", base64.URLEncoding.EncodeToString(buf.Bytes()))
	response, _ := client.Do(reqest)

	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)
	//body为 []byte类型
	returnbody, _ := ioutil.ReadAll(response.Body)
	status := response.StatusCode

	fmt.Println(reflect.TypeOf(returnbody))
	fmt.Println(status)
}
