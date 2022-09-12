// 동적계획법을 사용하면 더 빨리 처리할 수 있을 듯.
// 첫번째 8x8 만 전체 수행하고 나머지는 한줄씩만 처리하여 첫번째 결과에서 +-
package main

import (
	bio "bufio"
	. "fmt"
	"os"
)

var bw = "BWBWBWBW"
var wb = "WBWBWBWB"

func dc(s, d string) (count int) {
	for i, c := range s {
		if string(c) != d[i:i+1] {
			count++
		}
	}
	return
}

func main() {
	var n, m int
	Scan(&n, &m)

	r := bio.NewReader(os.Stdin)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		Fscanln(r, &s[i])
	}

	mn := 65
	s1, s2, t := 0, 0, 0
	for i := 0; i <= n-8; i++ {
		for j := 0; j <= m-8; j++ {
			s1, s2 = 0, 0
			for k := 0; k < 8; k++ {
				if k%2 == 0 {
					s1 += dc(s[i+k][j:j+8], wb)
					s2 += dc(s[i+k][j:j+8], bw)
				} else {
					s1 += dc(s[i+k][j:j+8], bw)
					s2 += dc(s[i+k][j:j+8], wb)
				}
			}
			if s1 <= s2 {
				t = s1
			} else {
				t = s2
			}
			if t < mn {
				mn = t
			}
		}
	}
	Println(mn)
}