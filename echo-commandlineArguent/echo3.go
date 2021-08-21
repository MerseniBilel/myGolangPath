package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t0 := time.Now()
	fmt.Println("The name of the program : ", os.Args[0])
	fmt.Println("args => ", strings.Join(os.Args[1:], " "))
	duration := time.Since(t0)
	println("program finished in ", duration)

	// index and value of each argument per line

	t0 = time.Now()
	for index, value := range os.Args[1:] {
		fmt.Println(index, "=>", value)
	}
	duration = time.Since(t0)

	println("program finished in ", duration)

}
