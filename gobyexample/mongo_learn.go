//教程参考 http://my.oschina.net/ffs/blog/300148
package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//注意User的字段首字母大写，不然不可见。
//通过bson:”name”这种方式可以定义MongoDB中集合的字段名，
//如果不定义，mgo自动把struct的字段名首字母小写作为集合的字段名
type Person struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`
}

type Command struct {
	Example string
	Type    string
	Intro   string
}

func main() {
	//注意在mongodb的配置文件中 /etc/mongod.conf 中 bindip 以及 port
	//要分别修改成  10.10.105.204 以及 27017
	//同服务器建立连接
	//或者不填写 bindip 这样就可以同时远程操控 或者用 mongoshell操控了 否则只能与固定的ip建立联系
	session, err := mgo.Dial("10.10.105.204:27018")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	//切换到对应的数据库
	c := session.DB("test").C("people")
	c2 := session.DB("test").C("Command")
	//插入对应的信息
	err = c.Insert(&Person{"Ale", "111111"}, &Person{"Cla", "222222222"})
	err = c2.Insert(&Command{"testexample", "testtype", "testintro"})

	if err != nil {
		panic(err)
	}
	result := Person{}
	resultb := Command{}

	//通过 name Ale 来定位到 对应行
	//格式是 bson 的格式 注意这里的关键字要采用小写的形式
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	err = c2.Find(bson.M{"type": "testtype"}).One(&resultb)

	if err != nil {
		panic(err)
	}
	fmt.Println("Phone:", result.Phone)
	fmt.Println("Example:", resultb.Example)

}
