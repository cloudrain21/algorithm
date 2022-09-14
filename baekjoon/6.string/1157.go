package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	. "strings"
)

func main() {
	var s string
	rd := bufio.NewReader(os.Stdin)
	Fscan(rd, &s)

	r := make([]int, 26)

	mx := 0
	mi := -1
	for i:=0; i<len(s); i++ {
		t := ToUpper(s[i:i+1])
		i := int(t[0]) - 'A'
		r[i]++
		if r[i] > mx {
			mx = r[i]
			mi = i
		}
	}

	sort.Ints(r)
	if r[25] == r[24] {
		Println("?")
	} else {
		Println(string('A' + mi))
	}
}
