//show some common using of the regular expression in golang
package BasicUsings

import (
	"bytes"
	"fmt"
	"regexp"
)

func RegxOperation() {
	//test weather a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peachabcde")
	fmt.Println(match)

	//use this to compile a pattern
	r, _ := regexp.Compile("p([a-z]+)ch")

	//see if match
	fmt.Println(r.MatchString("peach"))

	//find a string according the pattern from the source string
	fmt.Println(r.FindString("peach punch pinch"))

	//find the start and end index of first matching string in source
	fmt.Println("the index")
	fmt.Println(r.FindStringIndex("peach punch"))

	//find the whole matchstring and the submatches string
	//in this example the string maching p([a-z]+)ch and ([a-z]+) will be returned
	fmt.Println(r.FindStringSubmatch("peach punch"))

	//find index similar as FindStringIndex [start1 end1 start2 end2...]
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	//find all the substring (FindString only find the first matching substring)
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// simillar as the functions above
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	//find the first two submatch string
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	//the arhuments could be the form of []byte
	fmt.Println(r.Match([]byte("peach")))

	//when create a constant of regulare expression
	//you should use the mustcompile(which is different with compile function)
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	//replace the subset of the value by regular expression
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	//replacing?
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
