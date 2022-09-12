package main

import (
	. "bufio"
	"fmt"
	"os"
)

func main() {
	var n,c,sum int
	var s string
	fmt.Scan(&n)

	r := NewReader(os.Stdin)
	for i:=0; i<n; i++ {
		fmt.Fscanln(r, &s)
		c = 0
		sum = 0
		for j:=0; j<len(s); j++ {
			if "O" == s[j:j+1] {
				c++
			} else {
				c = 0
			}
			sum += c
		}
		fmt.Println(sum)
	}
}