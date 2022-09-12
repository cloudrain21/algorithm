package main

import (
	"bytes"
	"fmt"
)

var a [][]int

func p(x, y int, b bool) {
	if !b {
		return
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !(i == 1 && j == 1) {
				a[x+i][y+j] = 1
			}
		}
	}
}

func r(n, x, y int, b bool) {
	if n == 1 {
		p(x, y, b)
		return
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !(i == 1 && j == 1) {
				r(n/3, x+i*n, y+j*n, true)
			}
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	a = make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}

	r(n/3, 0, 0, true)

	var b bytes.Buffer
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			s := "*"
			if a[i][j] == 0 {
				s = " "
			}
			b.WriteString(s)
		}
		b.WriteString(string('\n'))
	}
	fmt.Print(b.String())
}
