package main

import (
	"bufio"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"os"
	"strings"
)

type Command struct {
	Type  string `bson:"type"`
	Intro string `bson:"intro"`
}

type Dockerfile struct {
	Template string
	Type     string
}

func main() {

	session, err := mgo.Dial("10.10.105.204:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	//切换到对应的数据库
	c := session.DB("toturial").C("cmds")

	//read from the file
	//fmt.Println(os.Getwd())
	f, err := os.Open("tutorial/tuto_a.md")

	if err != nil {
		panic(err)
	}
	readf := bufio.NewReader(f)

	Type := ""
	Intro := ""
	for {
		//读出内容保存为string 每次读到以'\n'为标记的位置
		line, err := readf.ReadString('\n')
		//fmt.Print(line)
		if err == io.EOF {
			break
		}

		if strings.Contains(line, "####") {

			if Type != "" && Intro != "" {
				fmt.Println("Type:", Type, "Intro:", Intro)
				err = c.Insert(&Command{Type, Intro})
				//清零
				Intro = ""
			}

			Type = strings.TrimLeft(line, "#### ")
			Type = strings.TrimRight(Type, "\n")
			fmt.Println(Type)
			continue

		}

		Intro = Intro + line

	}
	result := Command{}
	//注意存在mongodb中的 type 有一个\n
	err = c.Find(bson.M{"type": "FROM"}).One(&result)
	if err != nil {
		panic(err)
	}

}
