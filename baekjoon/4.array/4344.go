package main

import (
	bio "bufio"
	. "fmt"
	"os"
)

func main() {
	var tc, n, score int

	Scan(&tc)

	r := bio.NewReader(os.Stdin)
	for i:=0; i<tc; i++ {
		Fscan(r, &n)

		sum := 0.0
		a := make([]int,0)
		for j:=0; j<n; j++ {
			Fscan(r, &score)
			a = append(a, score)
			sum += float64(score)
		}

		avg := sum / float64(n)
		c := 0
		for i:=0; i<n; i++ {
			if float64(a[i]) > avg  {
				c++
			}
		}
		Printf("%.3f%s\n", float64(c*100)/float64(n), "%")
	}
}