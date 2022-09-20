package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	m := map[string]struct{}{}

	l := len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			b := i
			e := j
			m[s[b:e]] = struct{}{}
		}
	}
	fmt.Println(len(m))
}
