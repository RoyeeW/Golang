package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args[1:] {
		fmt.Println("key=", k, "|value=", v)
	}
}
