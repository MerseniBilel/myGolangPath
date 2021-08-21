package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// ? print the content of a URL

// TODO : use io.Copy instead of ioutil.ReadAll()
// TODO : auto add the http:// if missing
// TODO : print the http status code resp.Status
func main() {

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("there is an error ", err)
			os.Exit(1)
		}
		fmt.Println(" response code : ", resp.Status)
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println("error while reading the requests body ", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
