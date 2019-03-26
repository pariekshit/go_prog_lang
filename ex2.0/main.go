package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	read := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		fmt.Println(read)
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			fmt.Printf("HTTP Status is in the %d range", resp.StatusCode)
		} else {
			fmt.Println("Argh! Broken")
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
