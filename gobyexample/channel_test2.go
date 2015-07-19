package main

//参考go 并发编程实战 p337
import (
	"log"
	"strconv"
	"time"
)

type Person struct {
	name string
	age  int
	addr string
}

var oldpersonarray = [5]Person{}
var newpersonarray = [5]Person{}

type PersonHandler interface {
	Batch(origs <-chan Person) <-chan Person
	Handle(orig *Person)
}

//struct 实现了personhandler 接口
type PersonHandlerImpl struct{}

//从origs接收信息 处理之后再返回给新的channel
func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
	dests := make(chan Person, 100)
	go func() {
		for {
			p, ok := <-origs
			if !ok {
				close(dests)
				break
			}
			handler.Handle(&p)
			log.Printf("old value : %v\n", p)
			//time.Sleep(time.Second)
			dests <- p
		}
	}()
	return dests
}

//这里要使用引用传递
func (handler PersonHandlerImpl) Handle(orig *Person) {
	orig.addr = "new address"
}

func getPersonHandler() PersonHandler {
	return &PersonHandlerImpl{}

}

//print the oldpersonarray into the chan<-Person
func fetchPerson(origs chan<- Person) {
	for _, v := range oldpersonarray {
		//fmt.Printf("get the value : %v %v \n", k, v)
		time.Sleep(time.Second)
		origs <- v
	}
	close(origs)

}

//fetch the value from the channel and store it into the newpersonarray
func savePerson(dest <-chan Person) <-chan int {
	intChann := make(chan int)
	go func() {
		index := 0
		for {
			p, ok := <-dest
			if !ok {
				break
			}
			//time.Sleep(time.Second)
			log.Printf("new value transfer %v \n", p)

			newpersonarray[index] = p
			index++

		}

		intChann <- 1
	}()
	return intChann
}

func init() {
	//使用range的话是值传递 这里要给oldpersonarray赋值进来
	tmplen := len(oldpersonarray)
	for i := 0; i < tmplen; i++ {
		oldpersonarray[i].addr = "old address"
		oldpersonarray[i].age = i
		oldpersonarray[i].name = strconv.Itoa(i)

	}

	log.Printf("first print init value : %v\n", oldpersonarray)

}
func main() {

	handeler := getPersonHandler()
	origs := make(chan Person, 100)
	dests := handeler.Batch(origs)
	//go func() { fetchPerson(origs) }()
	// 不加go func的话 要等这句执行完 才能执行下一句
	// 则orgis信息都输出 完全关闭掉 这个时候 从dest接收信息的语句才开始执行
	// 所以不会动态输出 这句加上go func的话 就会没隔 1s 动态输出
	// 如果将fetchPerson 再往前面放一句 则old value也不会动态输出
	fetchPerson(origs)
	sign := savePerson(dests)
	<-sign
	log.Printf("last print new value : %v \n", newpersonarray)

}
