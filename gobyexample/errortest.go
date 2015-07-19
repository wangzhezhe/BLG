package main

import (
	"errors"
	"fmt"
	"net"
)

type Customerror struct {
	infoa string
	infob string
	Err   error
}

func (cerr Customerror) Error() string {
	errorinfo := fmt.Sprintf("infoa : %s , infob : %s , original err info : %s ", cerr.infoa, cerr.infob, cerr.Err.Error())
	return errorinfo
}

func main() {
	//方法一：
	//采用errors包的New方法 返回一个err的类型
	var err error = errors.New("this is a new error")
	//由于已经实现了error接口的方法 因此可以直接调用对应的方法
	fmt.Println(err.Error())

	//方法二：
	//采用fmt.Errof 将string信息转化为error信息 并返回
	err = fmt.Errorf("%s", "the error test for fmt.Errorf")
	fmt.Println(err.Error())

	//方法三：
	//采用自定义的方式实现一个error的 一个duck 类型
	err = &Customerror{
		infoa: "err info a",
		infob: "err info b",
		Err:   errors.New("test custom err"),
	}

	fmt.Println(err.Error())

	if cuserr, ok := err.(Customerror); ok {
		fmt.Println(cuserr.Error())
	} else {
		fmt.Println(err.Error())

	}
	err, _ = err.(net.Dial())

}
