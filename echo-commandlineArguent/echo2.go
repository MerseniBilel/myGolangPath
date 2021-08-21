package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	// * the range function
	// skip the value take the index
	// for i, _ := range [slice]

	// skip the index, take the value
	// for _, value := range [slice]

	// take the index and the value
	//for i, value := range [slice]

	// default : take the index
	// for [variable] := range [slice]

	for _, args := range os.Args[1:] {
		s += sep + args
		sep = " "
	}
	fmt.Println(s)
}
