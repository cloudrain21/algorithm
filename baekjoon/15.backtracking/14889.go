package main

import (
	"bufio"
	"fmt"
	"os"
)

var min int
var chk []bool
var ar [][]int

func main() {
	var n, t int

	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)

	chk = make([]bool, n)
	for i := 0; i < n; i++ {
		var in []int
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &t)
			in = append(in, t)
			min += t
		}
		fmt.Fscanln(r)
		ar = append(ar, in)
	}

	recur(n, 0, 0)
	fmt.Println(min)
}

func recur(n, curPos, curChked int) {
	if curChked == n/2 {
		d := diff1(n)
		if d < min {
			min = d
		}
	} else {
		for i := curPos; i < n; i++ {
			chk[i] = true
			recur(n, i+1, curChked+1)
			chk[i] = false
		}
	}
}

func diff1(n int) int {
	var sum1, sum2 = 0, 0
	team1 := make([]int, n/2)
	team2 := make([]int, n/2)

	x, y := 0, 0
	for i:=0; i<n; i++ {
		if chk[i] {
			team1[x] = i
			x++
		} else {
			team2[y] = i
			y++
		}
	}

	for i := 0; i < n/2; i++ {
		for j := i + 1; j < n/2; j++ {
			sum1 += ar[team1[i]][team1[j]] + ar[team1[j]][team1[i]]
			sum2 += ar[team2[i]][team2[j]] + ar[team2[j]][team2[i]]
		}
	}

	if sum1 > sum2 {
		return sum1 - sum2
	} else {
		return sum2 - sum1
	}
}