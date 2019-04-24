package main

import (
	"fmt"
)

func max(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	max := vals[0]
	for _, val := range vals[1:] {
		if max < val {
			max = val
		}
	}
	return max, true
}

func min(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	min := vals[0]
	for _, val := range vals[1:] {
		if min > val {
			min = val
		}
	}
	return min, true
}

func check(vals []int) {
	fmt.Println(vals)
	max, ok := max(vals...)
	if !ok {
		fmt.Println("Error has occurred.")
	} else {
		fmt.Println("Max:", max)
	}

	min, ok := min(vals...)
	if !ok {
		fmt.Println("Error has occurred.")
	} else {
		fmt.Println("Min:", min)
	}
}

func main() {
	check([]int{})
	check([]int{1, 2, 3})
	check([]int{-1, 5, 6, 7})
	check([]int{2, 4, 6, 8})
	check([]int{-2, -6, -10})
}
