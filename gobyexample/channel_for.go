package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	names := []string{"E", "H", "R", "J", "M", "N", "O", "P"}
	for _, name := range names {
		go func(who string) {
			time.Sleep(1000 * time.Nanosecond)
			fmt.Printf("Hello , %s \n", who)
		}(name)

	}
	runtime.Gosched()

}

/* output:
Hello , E
Hello , H
Hello , R
Hello , J
Hello , M
*/
