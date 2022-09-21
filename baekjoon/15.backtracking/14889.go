package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, t int
	var a [][]int

	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	for i:=0; i<n; i++ {
		var in []int
		for j:=0; j<n; j++ {
			fmt.Fscan(r, &t)
			in = append(in, t)
		}
		fmt.Fscanln(r)
		a = append(a, in)
	}

	// 최대 100 명을 50명씩 짝지었을 경우 전력 합산 diff의 최소값 구하기
}