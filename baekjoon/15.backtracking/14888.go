// 더 빠르고 짧은 코드로 수정 가능하다. 다시 해볼 것!!
package main

import (
	"fmt"
	"math"
)

const (
	PL = iota
	MI
	MU
	DI
)

var o []int
var of []bool
var na []int

var mx = -1000000001
var mn = 1000000001

func ca(x, y, o int) int {
	switch o {
	case PL:
		return x + y
	case MI:
		return x - y
	case MU:
		return x * y
	case DI:
		if x < 0 {
			return int(-(math.Abs(float64(x)) / float64(y)))
		}
		return x / y
	default:
		panic("invalid operator")
	}
}

func cc(arr []int, operStack []int) int {
	sm := arr[0]
	for i:=1; i<len(arr); i++ {
		sm = ca(sm, arr[i], operStack[i-1])
	}
	return sm
}

func ff(operStack []int) {
	if len(na)-1 == len(operStack) {
		r := cc(na, operStack)
		if r > mx {
			mx = r
		}
		if r < mn {
			mn = r
		}
		return
	}

	for i:=0; i<len(o); i++ {
		if !of[i] {
			of[i] = true
			operStack = append(operStack, o[i])
			ff(operStack)
			of[i] = false
			operStack = operStack[:len(operStack)-1]
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	o1 := make([]int, 4)
	na = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&na[i])
	}

	for i := 0; i < 4; i++ {
		fmt.Scan(&o1[i])
		for j := 0; j < o1[i]; j++ {
			o = append(o, i)
			of = append(of, false)
		}
	}

	var operStack []int
	ff(operStack)
	fmt.Println(mx)
	fmt.Println(mn)
}
