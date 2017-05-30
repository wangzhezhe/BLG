package reflection

import (
	"errors"
	"reflect"
)

var m = make(map[string]reflect.Type)

func init() {

	m["mysql"] = reflect.TypeOf(Mysql{})
	m["redis"] = reflect.TypeOf(Redis{})
	m["influx"] = reflect.TypeOf(Influx{})
}

func getInstance(name string) (interface{}, error) {
	t, ok := m[name]
	if ok {
		instance := reflect.New(t).Elem().Interface()
		return instance, nil
	} else {
		err := errors.New("type: " + name + " is not supported")
		return nil, err
	}

}
