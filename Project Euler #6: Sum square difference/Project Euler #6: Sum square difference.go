package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n int64
		fmt.Scanf("%d", &n)

		first := int64(math.Pow(float64(n*(n+1))/2, 2))
		second := int64((n * (n + 1) * (2*n + 1)) / 6)

		difference := first - second

		fmt.Printf("%d\n", difference)
	}
}
