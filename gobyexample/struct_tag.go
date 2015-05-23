// reference https://www.9696e.com/archives/1187

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	//貌似tag可以用""标记  也可以用``标记
	Name   string `user name`
	Passwd string `user password`
}

func main() {
	user := &User{"chronos", "pass"}

	//通过反射 获取type的定义 这里的s 与结构体里面的具体元素相关联了起来
	s := reflect.TypeOf(user).Elem()
	fmt.Println(reflect.TypeOf(s))

	//利用反射得到tag的相关信息
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag)
	}

}
