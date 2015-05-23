package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//change the lower case into the capital letter
func main() {

	//wrap the stdin using a buffered scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		//Text returns the current token, here the next line, from the input

		ucl := strings.ToUpper(scanner.Text())

		fmt.Println(ucl)

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}
