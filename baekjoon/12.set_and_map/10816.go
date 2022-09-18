package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	c := make(map[int]int)

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscan(r, &n)

	for i:=0; i<n; i++ {
		var tmp int
		fmt.Fscan(r, &tmp)
		if _, ok := c[tmp]; ok {
			c[tmp]++
		} else {
			c[tmp] = 1
		}
	}

	fmt.Fscan(r, &m)
	for i:=0; i<m; i++ {
		var tmp int
		fmt.Fscan(r, &tmp)
		if v, ok := c[tmp]; ok {
			tmp = v
		} else {
			tmp = 0
		}
		fmt.Fprintf(w, "%d ", tmp)
	}
	w.Flush()
}
