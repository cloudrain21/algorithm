package main

import . "fmt"

func main() {
	n := 0
	s := ""
	m := make(map[string]bool,0)

	Scan(&n)
	ct := 0
	for i:=0; i<n; i++ {
		Scan(&s)

		m = map[string]bool{}
		b := ""
		j := 0
		for j=0; j<len(s); j++ {
			t := s[j:j+1]
			if t != b {
				if ok := m[t]; ok {
					break
				}
				m[t] = true
				b = t
			}
		}
		if j >= len(s) {
			ct++
		}
	}
	Println(ct)
}
