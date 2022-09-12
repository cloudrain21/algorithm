package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	var n int

	r := bufio.NewReader(os.Stdin)

	Scan(&n)
	a := make([]int, n)
	for i:=0; i<n; i++ {
		Fscanln(r, &a[i])
	}

	sort.Ints(a)
	for i:=0; i<n; i++ {
		Println(a[i])
	}
}
