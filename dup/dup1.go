package main

import (
	"bufio"
	"fmt"
	"os"
)

// prints the text of each line that appearse
// more than once in standard input

func main() {
	// create a map with string keys and int values
	counts := make(map[string]int)

	// create a new scanner
	input := bufio.NewScanner(os.Stdin)

	// read values
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}

	// print the dublicated values
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
