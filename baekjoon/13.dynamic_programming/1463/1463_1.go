// 다른 답을 참고한 후 푼 방법. recursive 보다 성능이 훨씬 좋음.
package main

import (
	"fmt"
)

var dd []int

func mn(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func main() {
	var n int
	fmt.Scan(&n)
	dd = make([]int, n+1)

	dd[1] = 0
	for i:=2; i<=n; i++ {
		dd[i] = dd[i-1]
		if i%2 == 0 {
			dd[i] = mn(dd[i], dd[i/2])
		}
		if i%3 == 0 {
			dd[i] = mn(dd[i], dd[i/3])
		}
		dd[i]++
	}

	fmt.Println(dd[n])
}
