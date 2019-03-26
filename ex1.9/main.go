package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		start := time.Now()
		if strings.HasPrefix(url, "http") {
			fmt.Printf("Already has http in %s\n", url)
		} else {
			fmt.Printf("Didnt found http in %s. So added http://\n", url)
			url = "http://" + url
		}
		time.Sleep(10 * time.Second)
		resp, err := http.Get(url)
		fmt.Println(resp.Status)
		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			fmt.Println("HTTP Status is in the 2xx range")
		} else {
			fmt.Println("Argh! Broken")
		}
		if err != nil {
			fmt.Printf("Error occured while getting %s and error is %s", url, err)
		} else {
			fmt.Println("No Error Found ")
		}
		elapsed := time.Since(start)
		log.Printf("Time took %s", elapsed)
	}
}
