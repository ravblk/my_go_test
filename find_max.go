package main

import "fmt"

func find_max(x ...int) int {
	max := x[0]
	for _, per := range x {
		if max < per {
			max = per
		}
	}
	return max
}

func main() {

	fmt.Println(find_max(1, 2, 7, 2, 15, 34, 12, 12, 34, 45, 1, 12, 3, 54))
}