package main

import "fmt"

var mGrade = map[int]string {
	10: "A",
	9: "A",
	8: "B",
	7: "C",
	6: "D",
	5: "F",
	4: "F",
	3: "F",
	2: "F",
	1: "F",
	0: "F",
}

func main() {
	var score int

	for {
		fmt.Scanf("%d", &score)
		if score < 0 || score > 100 {
			continue
		}
		break
	}

	fmt.Println(mGrade[score/10])
}
