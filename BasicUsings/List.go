package BasicUsings

import (
	"container/list"
	"fmt"
	//"sort"
	"strings"
)

type Iners struct {
	Inera string
	Inerb string
}

type Element struct {
	Itema string
	Itemb string
	Itemc string
	Iners
}

func ListBasicOperation() {

	items := list.New()
	//items.MoveToFront()
	for _, x := range strings.Split("ABCDEFGH", "") {
		items.PushFront(x)
	}
	items.PushBack(9)

	element := &Element{
		Itema: "a",
		Itemb: "b",
		Itemc: "c",
		Iners: Iners{Inera: "d", Inerb: "e"},
	}
	items.PushBack(element)

	for element := items.Front(); element != nil; element = element.Next() {
		//attention the type of value in list is interface{}
		switch value := element.Value.(type) {
		case string:
			fmt.Printf("%s ", value)
		case int:
			fmt.Printf("%d ", value)
		default:
			fmt.Printf("%v ", value)
		}

	}

	//fmt.Println()
	//sort the list
	//sort.Search()
}
