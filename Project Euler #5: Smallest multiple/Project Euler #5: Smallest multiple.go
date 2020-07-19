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
		ans := 1
		for true {
			flag := 1
			for j := 1; j <= n; j++ {
				if ans%j != 0 {
					flag = 0
					ans += 1
					break
				}
			}
			if flag == 1 {
				fmt.Printf("%d\n", ans)
				break
			}
		}
	}
}
