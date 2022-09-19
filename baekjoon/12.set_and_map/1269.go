package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, t int

	r := bufio.NewReader(os.Stdin)
	m := make(map[int]struct{}, 0)

	fmt.Fscan(r, &a, &b)
	for i := 0; i < a; i++ {
		fmt.Fscan(r, &t)
		m[t] = struct{}{}
	}

	fmt.Fscanln(r)
	for i := 0; i < b; i++ {
		fmt.Fscan(r, &t)
		if _, ok := m[t]; ok {
			delete(m, t)
		} else {
			m[t] = struct{}{}
		}
	}

	fmt.Println(len(m))
}
