package BasicFunctions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func RequestDebug() {

	//start a basic server before testing
	resp, err := http.Get("http://127.0.0.1:8888/")

	//Attention to close the body reading !!!
	defer resp.Body.Close()
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	//copy the http response out
	//this dump operation do not affect the following funtion
	//similar function could be used in debugging request function by http.DumpRequest
	resa, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Printf("http debug : \n%s\n", resa)
	//do some operations for the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error %+v", err)
	}

	fmt.Printf("\nget response again:\n%+v,", string(body))
}
