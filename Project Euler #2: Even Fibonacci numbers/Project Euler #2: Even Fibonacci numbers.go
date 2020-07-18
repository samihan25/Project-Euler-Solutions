package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int64
		fmt.Scanf("%d", &n)

		var sum int64
		sum = 2

		var first int64
		first = 1
		var second int64
		second = 2

		for true {
			var temp int64
			temp = second + first
			first = second
			second = temp
			if second > n {
				break
			}
			if second%2 == 0 {
				sum = sum + second
			}
		}

		fmt.Printf("%d\n", sum)
	}
}
