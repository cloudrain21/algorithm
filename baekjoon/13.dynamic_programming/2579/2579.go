// 종료조건을 제대로 못찾았었음.
package main

import (
	"bufio"
	"fmt"
	"os"
)

var arr []int
var d []int
var mx = 0

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func m(n int) int {
	if n == 1 {
		return max(0, arr[1])
	} else if n == 2 {
		return max(max(arr[1], arr[2]), arr[2] + arr[1])
	} else {
		return max(arr[3] + arr[1], arr[2] + arr[3])
	}
}

func r(n int) int {
	if n <=3 {
		return m(n)
	}

	if d[n] > 0 {
		return d[n]
	}

	v1 := r(n - 2) + arr[n]
	v2 := r(n - 3) + arr[n-1] + arr[n]

	d[n] = max(v1, v2)

	return d[n]
}

func main() {
	var n int
	fmt.Scan(&n)

	rr := bufio.NewReader(os.Stdin)
	arr = make([]int, n+1)
	d = make([]int, n+1)

	for i:=1; i<=n; i++ {
		fmt.Fscan(rr, &arr[i])
	}

	fmt.Println(r(n))
}
