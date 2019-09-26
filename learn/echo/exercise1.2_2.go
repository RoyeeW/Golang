package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print("index:", i, "|value:", os.Args[i], "\n")
	}
}
