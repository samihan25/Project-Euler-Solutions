package main

import (
	"fmt"
	"math"
)

func is_prime(no int64) bool {
	var temp int64
	temp = int64(math.Sqrt(float64(no)))
	for i := int64(2); i <= temp; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int64
		fmt.Scanf("%d", &n)

		var flag int
		flag = 0

		var temp int64
		temp = int64(n / 2)

		for j := int64(2); j < temp; j++ {
			if n%j == 0 {
				var xyz int64
				xyz = int64(n / j)
				if is_prime(xyz) {
					fmt.Printf("%d\n", xyz)
					flag = 1
					break
				}
			}
		}
		if flag == 0 {
			fmt.Printf("%d\n", n)
		}
	}
}
