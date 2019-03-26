package main

import (
	"fmt"
	"io"
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
			fmt.Printf("HTTP Status is in the %d range ........... ", resp.StatusCode)
			fmt.Println("test")
		} else {
			fmt.Println("Argh! Broken")
		}
		if err != nil {
			fmt.Println("ERROR")

		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Println("test")

		}
	}
}
