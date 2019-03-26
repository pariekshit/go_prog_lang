package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main1() {
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
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}
		fmt.Println(os.Stdout)
	}
}
