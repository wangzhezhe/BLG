package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		//errors.New construct a basic error with given info
		return -1, errors.New("can't work with 42")
	}
	//the nil is the error message indicate there are no error message
	return arg + 3, nil
}

//custome the error
//It’s possible to use custom types
//as errors by implementing the Error() method on them
//first to customize your own struct argError
//than implementing the metods on your struct(struct type is receiver)
//when return the error using &structname{para1,para2...}to create the new struct
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f2(i);
		//adjust if the error info is nil
		e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}

	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	err := errors.New("this is an error")
	fmt.Println(err)
}
