package main

import "fmt"

func findMax(x ...int) int {
	max := x[0]
	for _, per := range x {
		if max < per {
			max = per
		}
	}
	return max
}

func main() {

	fmt.Println(findMax(1, 2, 7, 2, 15, 34, 12, 12, 34, 45, 1, 12, 3, 54))
}