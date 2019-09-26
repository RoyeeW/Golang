package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("start:", start)
	var s string
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i]
	}
	s = ""
	fmt.Printf("end: %.3fs", time.Since(start).Seconds())
	fmt.Println("\n---------------------------------------------------------")
	start2 := time.Now()
	fmt.Println(start2)
	strings.Join(os.Args[1:], "")
	fmt.Printf("end: %.3fs", time.Since(start2).Seconds())
}
