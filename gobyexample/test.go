package main

import (
	"encoding/json"
	"fmt"
	//	"os"
)

func main() {
	image := []string{"test1", "test2"}
	image_json, temp := json.Marshal(image)
	fmt.Println(image_json, temp)

}
