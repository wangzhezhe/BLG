package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	//	"os"
)

func main() {
	image := []string{"test1", "test2"}
	image_json, temp := json.Marshal(image)
	fmt.Println(image_json, temp)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("wangzhezhe:wangzhezhe")))

}
