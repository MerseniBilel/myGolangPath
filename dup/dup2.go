package main

import (
	"bufio"
	"fmt"
	"os"
)

// prints the count and text of lines that appear
// more than once in the input.
// It reads from stdin or from a list of named files.

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("no file loaded")
		os.Exit(0)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if errFound(err) {
				fmt.Fprintf(os.Stderr, "Dup2 %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d ===> %s \n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func errFound(err error) bool {
	return err != nil
}
