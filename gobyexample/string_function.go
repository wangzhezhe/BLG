//show some basic usage of strings(util package)
package main

import (
	"fmt"
	"reflect"
	s "strings"
)

//为Println指定了别名 操作更简单了
var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count:    ", s.Count("ttest", "t"))

	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	//find location of e in test
	p("Index:    ", s.Index("test", "e"))
	p("Join:     ", s.Join([]string{"a", "b"}, "-"))
	//将字符串a 重复输出5次
	p("Repeat:   ", s.Repeat("abc", 5))
	//将 fooo 中出现的o 替换为 0
	//最后一个参数为-1表示全都替换 要是大于0 表示前n个需要替换
	p("Replace:  ", s.Replace("fooo", "o", "0", -1))
	p("Replace:  ", s.Replace("fooo", "o", "0", 1))
	p("Replace:  ", s.Replace("fooo", "o", "0", 2))
	p("Replace:  ", s.Replace("fooo", "o", "0", 3))
	//按照 - 将字符串进行分割 返回的类型是 []string
	var str = s.Split("a-b-c-d-e", "-")
	p(reflect.TypeOf(str))
	p("Split:    ", str)
	p("ToLower:  ", s.ToLower("TEST"))
	p("ToUpper:  ", s.ToUpper("test"))

	p("Length:   ", len("ttest"))
	//取得 字符串 hello 的第一个字符 默认的 返回的是一个asc码
	//除非加一个string的强制转换 才能按照字符串的形式 输出
	p("Char:     ", "hello"[0])
	p("Char:     ", string("hello"[0]))

}
