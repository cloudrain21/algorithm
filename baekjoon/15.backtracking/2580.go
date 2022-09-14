package main

import (
	"fmt"
)

type loc struct {
	x int
	y int
}

var e []loc

func c1(a *[9][9]int, x,y int) {
	pos := 0
	v := 0
	sm := 0
	cnt := 0
	for i:=0; i<9; i++ {
		if a[x][i] > 0 {
			cnt++
			sm += a[x][i]
		} else {
			pos = i
		}
	}
	if cnt == 8 {
		v = 45 - sm
		a[x][pos] = v
	}

	pos = 0
	v = 0
	sm = 0
	cnt = 0
	for i:=0; i<9; i++ {
		if a[i][y] > 0 {
			cnt++
			sm += a[i][y]
		} else {
			pos = i
		}
	}
	if cnt == 8 {
		v = 45 - sm
		a[pos][y] = v
	}
}

func c2(a *[9][9]int, x, y int) {
	posx, posy := 0, 0
	v := 0
	sm := 0
	cnt := 0
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			if a[x+i][y+j] > 0 {
				cnt++
				sm += a[x+i][y+j]
			} else {
				posx, posy = x+i, y+j
			}
		}
	}

	if cnt == 8 {
		v = 45 - sm
		a[posx][posy] = v
	}
}

func main() {
	a := [9][9]int{}

	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			fmt.Scan(&a[i][j])
			if a[i][j] == 0 {
				e = append(e, loc{i,j})
			}
		}
	}

	for len(e) > 0 {
	}

	for i:=0; i<9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}
