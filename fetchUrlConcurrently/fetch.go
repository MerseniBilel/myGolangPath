package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	table := make(map[string]int64)
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch, table) // start a go routine using the go key
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2f's \n", secs)

	printmax(table, os.Args[1])
}

func printmax(table map[string]int64, url string) {
	max := table[url]
	maxindex := url
	for index, value := range table {
		if value > max {
			max = value
			maxindex = index
		}
	}
	fmt.Printf("the largest website is\n %s --> %7d bytes\n", maxindex, max)
}

func fetch(url string, ch chan<- string, table map[string]int64) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("while getting the %s : %v ", url, err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}
	table[url] = nbytes

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs   %7d    %s ", secs, nbytes, url)
}
