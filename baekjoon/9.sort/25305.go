package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	var n,k int

	Scan(&n, &k)

	r := bufio.NewReader(os.Stdin)

	a := make([]int, n)
	for i:=0; i<n; i++ {
		Fscan(r, &a[i])
	}

	sort.Ints(a)
	Println(a[n-k])
}
