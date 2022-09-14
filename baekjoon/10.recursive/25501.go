package main

import (
	"bufio"
	"fmt"
	"os"
)

var rcnt int

func main() {
	var n int
	fmt.Scan(&n)

	var s string
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s)
		rcnt = 0
		fmt.Println(isPalindrome(s), rcnt)
	}
}

func recursive(s string, l, r int) int {
	rcnt++
	if l >= r {
		return 1
	}

	if s[l] != s[r] {
		return 0
	}

	return recursive(s, l+1, r-1)
}

func isPalindrome(s string) int {
	return recursive(s, 0, len(s)-1)
}
