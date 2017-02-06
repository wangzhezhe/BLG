package BasicUsings

/*
import (
	"fmt"
	"reflect"
	"strconv"
)

type Binary uint64
type myInt int

type Stringer interface {
	String() string
}

func (i Binary) String() string {
	return strconv.Itoa(i.Get())
}

func (i Binary) Get() int {
	return int(i)
}
func (i myInt) Get() int {
	return int(i)
}

func (i myInt) String() string {
	return strconv.Itoa(i.Get())
}

func main() {
	b := Binary(20)
	s := Stringer(b)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(new(Stringer)))
	fmt.Println(reflect.ValueOf(s))

	var i interface{}
	fmt.Println(reflect.TypeOf(i))
	i = b
	fmt.Println(reflect.TypeOf(i).Implements(reflect.TypeOf(new(Stringer)).Elem()))
	fmt.Println(reflect.ValueOf(i))
	m := myInt(30)
	i = m
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))
	fmt.Println(reflect.ValueOf(interface{}(s)))
	reflect.New()
}
*/
