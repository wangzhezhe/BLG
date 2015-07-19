package main

//参考go 并发编程实战 p337
import (
	"fmt"
	"time"
)

type person struct {
	name string
	age  int
	addr string
}

type PersonHandler interface{
	Batch(origs<-chan Person) <-Person
	Handle(orig Person)
}


type PersonHandlerImpl struct{}

func (handler PersonHandlerImpl) Batch(origs <-chan )<-Person{
	
}

func (handler PersonHandlerImpl) Handle(orig Person){
	
}

func getPersonHandler() PersonHandler{
	
}

func fetchPerson(origis chan<-Person){
	
}

func savePerson(dest <-chan Person)<-byte{
	
}

func main() {
	
	handeler:=getPersonHandler()
	origs:=make(chan Person,100)
	dests:=handeler.Batch(origs)
	fetchPerson(origs)
	sign:=savePerson(dests)
    <-sign
}
