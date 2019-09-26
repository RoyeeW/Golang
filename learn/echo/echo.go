package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	var i int
	for i = 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		s = s + "|"
	}
	fmt.Println(os.Args)
	fmt.Println(s)
}
