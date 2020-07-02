package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	//"io/ioutil"   //you can use this package if you are using Readall instaed of copy
)

func main() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http")) {
			url = "http://" + url
		}
		r, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, r.Body)
		
		//b, err := ioutil.Readall(r.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Reading %s: %v\n", url, err)
			os.Exit(1)
		}
		s := r.Status
		r.Body.Close()
		fmt.Printf("%s\n\n\nResponsecode: %s\n", b,s)
	}
}
