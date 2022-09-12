package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	no int
	left *node
	right *node
}

var nm [][]node
var d [][]int
var f [][]int
var mx int
var n int

func r(a int) {
	if a == n {
		// get sum
	}
}

func main() {
	fmt.Scan(&n)

	r := bufio.NewReader(os.Stdin)

	nm = make([][]node, n)
	d = make([][]int, n)
	f = make([][]int, n)

	for i:=0; i<n-1; i++ {
		nm[i] = make([]node, i+1)
		d[i] = make([]int, i+1)
		f[i] = make([]int, i+1)
		for j:=0; j<i+1; j++ {
			fmt.Fscan(r, &nm[i][j].no)
			nm[i][j].left = &nm[i+1][j]
			nm[i][j].right = &nm[i+1][j+1]
		}
	}
}
