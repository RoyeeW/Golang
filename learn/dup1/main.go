package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++ //输入的内容相同，count机会计数并自增1
		//fmt.Printf("%s\n", input.Text())
		//fmt.Println(counts)
	}
	fmt.Println("executed....")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d | %s\n", n, line)
		}
	}
}
