package main

import "fmt"

var n int
var a7 [][]int

func calc(x, y int, c int) {
	for i:=0; i<n; i++ {
		a7[x][i] += c
		a7[i][y] += c
	}

	a7[x][y] -= c

	for i:=1; i<n; i++ {
		x1, y1 := x-i, y-i
		if x1 >=0 && y1 >= 0 {
			a7[x1][y1] += c
		}

		x1, y1 = x+i, y+i
		if x1 < n && y1 < n {
			a7[x1][y1] += c
		}

		x1, y1 = x-i, y+i
		if x1 >= 0 && y1 < n {
			a7[x1][y1] += c
		}

		x1, y1 = x+i, y-i
		if x1 < n && y1 >= 0 {
			a7[x1][y1] += c
		}
	}
}

func main() {
	fmt.Scan(&n)

	a7 = make([][]int, n)
	for i:=0; i<n; i++ {
		a7[i] = make([]int, n)
	}

	calc(5,5,1)
	calc(5,5,-1)
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			fmt.Printf("%3d ", a7[i][j])
		}
		fmt.Println()
	}
}
