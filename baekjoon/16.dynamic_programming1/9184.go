package main

import (
	"bufio"
	"fmt"
	"os"
)

type st struct {
	a int
	b int
	c int
}

var m = map[st]int{}

func w(a,b,c int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}

	if v, ok := m[st{a,b,c}]; ok {
		return v
	}

	if a > 20 || b > 20 || c > 20 {
		m[st{a,b,c}] = w(20,20,20)
	} else if a < b && b < c {
		m[st{a,b,c}] = w(a, b, c-1) + w(a, b-1, c-1) - w(a, b-1, c)
	} else {
		m[st{a,b,c}] = w(a-1, b, c) + w(a-1, b-1, c) + w(a-1, b, c-1) - w(a-1, b-1, c-1)
	}

	return m[st{a,b,c}]
}

func main() {
	var a, b, c int
	r := bufio.NewReader(os.Stdin)
	wter := bufio.NewWriter(os.Stdout)

	for {
		fmt.Fscan(r, &a, &b, &c)
		if a == -1 && b == -1 && c == -1 {
			break
		}
		fmt.Fprintf(wter, "w(%d, %d, %d) = %d\n", a, b, c, w(a,b,c))
	}
	wter.Flush()
}
