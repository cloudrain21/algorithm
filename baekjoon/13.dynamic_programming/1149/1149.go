package main

import (
	"bufio"
	"fmt"
	"os"
)

type color int
const (
	RED color = iota
	GRE
	BLU
)

var g [][3]int
var d [1001][3]int

func r(n int, c color) int {
	if n == 1 {
		return g[n][c]
	}

	if d[n][c] > 0 {
		return d[n][c]
	}

	nextColor := (c+1)%3
	r1 := r(n-1, nextColor)

	nextColor = (c+2)%3
	r2 := r(n-1, nextColor)

	s1 := r1 + g[n][c]
	s2 := r2 + g[n][c]

	if s1 < s2 {
		d[n][c] = s1
	} else {
		d[n][c] = s2
	}
	return d[n][c]
}

func main() {
	var n int
	fmt.Scan(&n)

	rr := bufio.NewReader(os.Stdin)

	g = make([][3]int, n+1)
	for i:=1; i<=n; i++ {
		fmt.Fscan(rr, &g[i][RED], &g[i][GRE], &g[i][BLU])
	}

	mn := 1000001
	for i:=0; i<3; i++ {
		d = [1001][3]int{}
		v := r(n, color(i))
		if v < mn {
			mn = v
		}
	}
	fmt.Println(mn)
}