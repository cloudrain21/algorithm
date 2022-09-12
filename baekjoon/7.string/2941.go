package main

import "fmt"


var ar = []string {"c=","c-","dz=","d-","lj","nj","s=","z="}

func main() {
	var s string
	var m = map[string]bool{}

	for _, v := range ar {
		m[v] = true
	}

	fmt.Scan(&s)

	cnt := 0
	for i:=0; i<len(s); i++ {
		j := i
		for ; j<len(s); j++ {
			t := s[i:j+1]
			if ok := m[t]; ok {
				cnt++
				i = j
				break
			}
		}
		if j >= len(s) {
			cnt++
		}
	}
	fmt.Println(cnt)
}
