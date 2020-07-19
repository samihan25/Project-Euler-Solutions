package main

import (
	"fmt"
	"math"
)

func is_palindrome(no int) bool {
	length := int(math.Log10(float64(no)) + 1)
	//var arr [length]int
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		arr[i] = no % 10
		no = int(no / 10)
	}

	for i := 0; i < int(length/2); i++ {
		if arr[i] != arr[length-i-1] {
			return false
		}
	}
	return true
	/*
		for no > 9 {
			left := int(no / int(math.Pow10(length-1)))
			right := int(no % 10)
			if left != right {
				return false
			}
			no = int((no % int(math.Pow10(length-1))) / 10)
			length = length - 2
		}
		return true
	*/
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for iii := 0; iii < t; iii++ {
		var n int
		fmt.Scanf("%d", &n)

		var max int
		max = 10201

		//found := 0

		for i := 101; i < 1000; i++ {
			for j := 101; j < 1000; j++ {
				temp := i * j
				if temp < n {
					if temp > max {
						if is_palindrome(temp) {
							max = temp
						}
					}
				} /*else {
					found = 1
					break
				}
				*/
			}
			/*
				if found == 1 {
					break
				}
			*/
		}
		fmt.Printf("%d\n", max)
	}
}
