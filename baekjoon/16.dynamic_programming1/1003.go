// 더 간단한 방법이 있다.
package main

import (
	"bufio"
	"fmt"
	"os"
)

var d [41]int

func fibo(n int) int {
	if n == 0 {
		d[0] = 1
		return 0
	}
	if n == 1 {
		d[1] = 1
		return 1
	}
	if d[n] <= 0 {
		d[n] = fibo(n-1) +fibo(n-2)
	}
	return d[n]
}

func main() {
	var n,t int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Scan(&n)

	for i:=0; i<n; i++ {
		fmt.Fscan(r, &t)
		fibo(t)
		if t == 0 {
			fmt.Fprintln(w, 1, 0)
		} else if t == 1 {
			fmt.Fprintln(w, 0, 1)
		} else {
			fmt.Fprintln(w, d[t-1], d[t])
		}
		w.Flush()
	}
}
