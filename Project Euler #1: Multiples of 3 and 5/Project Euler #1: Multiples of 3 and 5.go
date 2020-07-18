package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d", &n)

		var sum int64
		sum = 0

		for j := 1; j < n; j++ {
			if j%3 == 0 {
				sum = sum + int64(j)
			} else if j%5 == 0 {
				sum = sum + int64(j)
			}
		}
		fmt.Printf("%d\n", sum)
	}
}
