package main

import . "fmt"

func okay(a int) int {
	s := Sprintf("%d", a)
	b1 := int(s[1] - s[0])
	for i:=2; i<len(s); i++ {
		b2 := int(s[i]-s[i-1])
		if b1 != b2 {
			return 0
		}
		b1 = b2
	}
	return 1
}

func main() {
	var n,s int

	Scan(&n)
	if n < 100 {
		s = n
	} else {
		for i:=100; i<=n; i++ {
			s += okay(i)
		}
		s += 99
	}
	Println(s)
}
