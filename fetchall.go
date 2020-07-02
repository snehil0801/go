package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	"time"
	"sync"
)
var wg sync.WaitGroup

func fetch(url string, ch chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nByte, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error while reading: %s\t%v\n", url, err)
		return
	}
	r.Body.Close()
	sec := time.Since(start).Seconds()
	ch <- fmt.Sprintf("\n\n\n====== Elapsed time: %.2f while reding url: %s with respnse ===\n\n\n%s", sec, url, nByte)

}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	f, err := os.Create("output")
	if err != nil {
		fmt.Printf("failed in creation of file %v", err)
		f.Close()
		os.Exit(1)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		fmt.Fprintln(f, <-ch)

	}
	f.Close()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
