package main

import (
	"bufio"
	"fmt"
	"os"
)

var dd = [101]int{0,1,1,1}

func init() {
	for i:=4; i<=100; i++ {
		dd[i] = dd[i-2] + dd[i-3]
	}
}

func main() {
	var t, n int
	fmt.Scan(&t)

	rr := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	for i:=0; i<t; i++ {
		fmt.Fscan(rr, &n)
		fmt.Fprintln(w, dd[n])
	}
	w.Flush()
}
