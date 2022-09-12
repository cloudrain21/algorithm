package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	mx := 0
	ix := 0
	for i:=1; i<=9; i++ {
		var c int
		fmt.Fscanln(r, &c)
		if c > mx {
			mx = c
			ix = i
		}
	}
	fmt.Println(mx)
	fmt.Println(ix)
}