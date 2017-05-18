package store

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/wangzhezhe/BLG/Components/StateMachine/util/queue"
	"github.com/wangzhezhe/BLG/Components/StateMachine/util/wait"
)

const (
	statea = iota
	stateb
	statec
	stated
	statee
)

const (
	objectNum       = 3
	defaultStateNum = 5
)

var StateStore map[string]int
var ObjSlice []string
var workQueue = queue.New()

func Start() {
	//init state store
	StateStore := make(map[string]int)
	ObjSlice = []string{"obja", "objb", "objc"}
	StateStore[ObjSlice[0]] = 0
	StateStore[ObjSlice[1]] = 1
	StateStore[ObjSlice[2]] = 2

	ticker := time.NewTicker(time.Second * 3)

	//start to create the event
	for _ = range ticker.C {

		newObjIndex := rand.Intn(objectNum)
		newState := rand.Intn(defaultStateNum)
		//put the event into the queue if the state change
		var event string
		if StateStore[ObjSlice[newObjIndex]] != newState {
			StateStore[ObjSlice[newObjIndex]] = newState
			event = fmt.Sprintf("change the index %s's state into %d at time %s", ObjSlice[newObjIndex], newState, time.Now().String())

		}
		//fmt.Printf("create the event: %s\n", event)
		workQueue.Add(event)
		//item, ifShutdown := workQueue.Get()
		//fmt.Println(item, ifShutdown)
	}

}

func getQueue() {
	var item interface{}
	//fmt.Println(workQueue.Len())
	item, ifShutdown := workQueue.Get()
	if item.(string) != "" {
		fmt.Printf("get event from the queue: %+v, queue status %+v\n", item, ifShutdown)
	}

	//do specific operaions for every event ...

}

func Watch() {

	signal := make(chan struct{})
	//after closing the signal
	//value, ok := signal will return {} ,false
	defer close(signal)
	go wait.Until(getQueue, time.Second, signal)
	time.Sleep(time.Second * 15)
	fmt.Println("time out")
	//data := struct{}{}
	//signal <- data
	//value, ok := <-signal

	//fmt.Println("channel value %+v channel boolean %+v", value, ok)

}
