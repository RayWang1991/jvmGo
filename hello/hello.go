package main

import (
	"fmt"
	"math"
)

func main() {
	sayhello()
	fmt.Println(max(1, 2))
	num1, num2, num3 := 1, 2, 3
	max, min := max_min_of_three(num1, num2, num3)
	fmt.Printf("for %d %d %d, max:%d, min:%d\n", num1, num2, num3, max, min)
}

func max(num1 int, num2 int) int {
	res := math.MaxInt32
	if num1 > num2 {
		res = num1
	} else {
		res = num2
	}
	return res
}

func min(num1 int, num2 int) int {
	res := math.MaxInt32
	if num1 < num2 {
		res = num1
	} else {
		res = num2
	}
	return res
}

func max_min_of_three(num1, num2, num3 int) (int, int) {
	max, min := max(max(num1, num2), num3), min(min(num1, num2), num3)
	return max, min
}
