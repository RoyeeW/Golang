package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	if len(os.Args) > 1 {
		for i := 0; i < len(os.Args); i++ {
			s = s + os.Args[i] + "|"
		}
		fmt.Println(s)
	} else {
		fmt.Println("Usage exercise.exe parm1 parm2 ....")
	}

}
