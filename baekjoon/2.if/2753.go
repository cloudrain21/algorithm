package main

import "fmt"

func isYun(year int) int {
	if year%4 != 0 {
		return 0
	}

	if year%100 != 0 {
		return 1
	}

	if year%400 == 0 {
		return 1
	}
	return 0
}

func main() {
	var year int

	for {
		fmt.Scanf("%d", &year)
		if year < 1 || year > 4000 {
			continue
		}
		break
	}

	fmt.Println(isYun(year))
}
