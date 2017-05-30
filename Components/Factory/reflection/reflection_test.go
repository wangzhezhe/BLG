package reflection

import (
	"fmt"
	"reflect"
	//"strings"
	"testing"
)

func TestGetInstance(t *testing.T) {

	strlist := []string{"mysql", "redis", "influx", "others"}
	for _, item := range strlist {
		instance, err := getInstance(item)
		if err != nil {
			t.Error(err)
			return
		}
		mapv, ok := m[item]
		if !ok {
			t.Error(err)
			return
		}

		fmt.Println(mapv.String())
		fmt.Println(reflect.TypeOf(instance).String())
		//compare reflect type directly
		if mapv == reflect.TypeOf(instance) {
			fmt.Println("current instance type is " + mapv.Name())
		} else {
			t.Error("get error type")
		}

		/*
			//transfer into the string and compare directly
			if strings.EqualFold(mapv.String(), reflect.TypeOf(instance).String()) {
				fmt.Println("current instance type is " + mapv.Name())
			} else {
				t.Error("get error type")
			}
		*/

	}
}
