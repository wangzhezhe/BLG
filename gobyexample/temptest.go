package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Env []string

func (env *Env) Map() map[string]string {
	if len(*env) == 0 {
		return nil
	}
	m := make(map[string]string)
	for _, kv := range *env {
		parts := strings.SplitN(kv, "=", 2)
		m[parts[0]] = parts[1]
	}
	return m
}
func main() {
	fmt.Println("print str '\"'")
	AuthConfig := map[string]string{
		"username": "w_hessen@126.com",
		"password": "3.1415",
		"email":    "w_hessen@126.com"}

	fmt.Println(AuthConfig)
	json_str, _ := json.Marshal(AuthConfig)
	fmt.Println(string(json_str))
	encode := base64.NewEncoding(string(json_str))
	fmt.Println(encode)

	env := Env{"a=b", "c=d"}
	m := env.Map()
	fmt.Println(m["a"])
	fmt.Println(m["c"])

}
