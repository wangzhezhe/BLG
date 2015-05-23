package main

import "os"

/*

A common use of panic is to abort
if a function returns an error value
that we don’t know how to (or want to) handle.
Here’s an example of panicking
if we get an unexpected error
when creating a new file


*/

func main() {
	panic("a problem")

	/*

	   in Go it is idiomatic to use error-indicating
	   return values wherever possible


	*/
	_, err := os.Create("/tmp/file")

	if err != nil {
		panic(err)
	}

}
