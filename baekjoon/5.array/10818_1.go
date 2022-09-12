package main

import (
	bi "bufio"
	. "fmt"
	"os"
)

func main() {
	t := 0
	Scanf("%d", &t)

	x := -1000001
	n := 1000001

	r := bi.NewReader(os.Stdin)
	for i:=0; i<t; i++ {
		var c int
		Fscan(r, &c)
		if c > x {
			x = c
		}
		if c < n {
			n = c
		}
	}

	Println(n, x)
}
