// 시간 초과
package main

import (
	"bufio"
	"fmt"
	"os"
)

var k, result int

func merge_sort(a []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		merge_sort(a, p, q)
		merge_sort(a, q+1, r)
		merge(a, p, q, r)
	}
}

func merge(a []int, p, q, r int) {
	i := p
	j := q + 1
	t := 0
	tmp := make([]int, len(a))
	for i <= q && j <= r {
		if a[i] <= a[j] {
			tmp[t] = a[i]
			i++
		} else {
			tmp[t] = a[j]
			j++
		}
		t++
	}

	for i <= q {
		tmp[t] = a[i]
		t++; i++
	}
	for j <= r {
		tmp[t] = a[j]
		t++; j++
	}
	i = p
	t = 0
	for i <= r {
		a[i] = tmp[t]
		k--
		if k == 0 {
			result = a[i]
		}
		i++; t++
	}
}

func main() {
	var n int
	fmt.Scan(&n, &k)

	r := bufio.NewReader(os.Stdin)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	result = -1
	merge_sort(a, 0, len(a)-1)
	fmt.Println(result)
}
