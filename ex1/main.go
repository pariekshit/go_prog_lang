package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("index : %d and value is: %s\n", i, os.Args[i])
	}
}
