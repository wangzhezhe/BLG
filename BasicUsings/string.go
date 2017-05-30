package BasicUsings

import (
	"fmt"
	"reflect"
	s "strings"
)

var str = s.Split("a-b-c-d-e", "-")

//using alias to simplify the operation
var p = fmt.Println

func split(r rune) bool {
	if r == 'a' {
		return true
	}
	return false
}

type point struct {
	x, y int
}

func StringOperations() {
	//create the new string \n was the ESC in string
	s1 := "hello\nworld!"
	//use `` to show the raw string
	s2 := `hello\nworld!`
	fmt.Println(s1)
	fmt.Println(s2)

	//formatting the value
	pt := point{1, 2}
	//%v print go value
	fmt.Printf("%v\n", pt)
	//If the value is a struct,
	//the %+v variant will include the struct’s field names.
	fmt.Printf("%+v\n", pt)
	//The %#v variant prints a Go syntax representation of the value,
	//i.e. the source code snippet that would produce that value.
	fmt.Printf("%#v\n", pt)

	//print the type of value
	//it is more convinient than using reflect
	fmt.Printf("%T\n", pt)

	//use sprintf to get a formateed string
	rets := fmt.Sprintf("return a string %d", 123)
	fmt.Println(rets)

	p("Contains: ", s.Contains("test", "es"))
	p("Count:    ", s.Count("ttest", "t"))

	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	//find location of e in test
	p("Index:    ", s.Index("test", "e"))
	p("Join:     ", s.Join([]string{"a", "b"}, "-"))
	//repeat the string for 5 times
	p("Repeat:   ", s.Repeat("abc", 5))
	//replace the "o" into "0"
	//if the last parameter is -1 , it means to replace all the "o"
	//if the last parameter is interger n , it means to replace first n "o" in string
	p("Replace:  ", s.Replace("fooo", "o", "0", -1))
	p("Replace:  ", s.Replace("fooo", "o", "0", 1))
	p("Replace:  ", s.Replace("fooo", "o", "0", 2))
	p("Replace:  ", s.Replace("fooo", "o", "0", 3))
	//split the string according to the "-" return type is []string

	p(reflect.TypeOf(str))
	p("Split:    ", str)
	//transfer between uppercase and lowercase
	p("ToLower:  ", s.ToLower("TEST"))
	p("ToUpper:  ", s.ToUpper("test"))

	p("Length:   ", len("ttest"))

	//return the first char in the string , defaut type is asc code
	p("Char:     ", "hello"[0])
	p("Char:     ", string("hello"[0]))

	//use trim to delete te extra character in string
	//invoke the trim right & trim left specifically
	p("trim", s.Trim("\n\n\naaaaaabcd-efgh-aaabcdaaaaa\n\n", "\n"))
	p("trimleft", s.TrimLeft("aaaaaabcd-efaaagh-aaabcdaaaaa", "a"))
	p("trimright", s.TrimRight("aaaaaabcd-eaaafgh-aaabcdaaaaa", "a"))

	//Attention!!! all the characters in second string will be deleted
	p("trimright", s.TrimRight("apm1-1.1.1", "-1.1"))

	//only trim the prefix and the suffix string
	p("trimprefix", s.TrimPrefix("aaaaaabcd-efaaagh-aaabcdaaaaa", "aaaaaabcd"))
	p("trimsuffix", s.TrimSuffix("aaaaaabcd-efaaagh-aaabcdaaaaa", "aaabcdaaaaa"))

	//find the index of the character in string, the first index is 0
	p("index:", s.Index("sdfgh", "d"))

	//using function to locate the index of the first character
	p("index:", s.IndexFunc("nihaoma", split)) //3

	//Attention!!! using equalfold function to adjust the string equal instead of "=="
	p("abc==def?", s.EqualFold("abc", "def"))

	p("abc==abc?", s.EqualFold("abc", "abc"))
}
