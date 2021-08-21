package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		// using ioutil u dont have to close the file
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d ===> %s\n", n, line)
		}
	}
}
