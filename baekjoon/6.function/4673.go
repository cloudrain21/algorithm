package main

import "fmt"

func ss(a int) int {
	s := 0
	for _, c := range fmt.Sprintf("%d", a) {
		s += int(c - '0')
	}
	return s
}

func main() {
	a := make([]bool, 10001)

	for i:=1; i<len(a); i++ {
		if ! a[i] {
			fmt.Println(i)
		}
		r := i + ss(i)
		if r <= 10000 {
			a[r] = true
		}
	}
}