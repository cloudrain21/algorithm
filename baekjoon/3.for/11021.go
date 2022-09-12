package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var total int
	type caseNum struct {
		a int
		b int
	}
	var nums []caseNum

	fmt.Scanf("%d", &total)

	r := bufio.NewReader(os.Stdin)
	for i := 0; i < total; i++ {
		var a, b int
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		nums = append(nums, caseNum{a, b})
	}

	var b bytes.Buffer
	for i := 0; i < total; i++ {
		s := fmt.Sprintf("Case #%d: %d\n", i+1, nums[i].a+nums[i].b)
		b.WriteString(s)
	}

	fmt.Print(b.String())
}