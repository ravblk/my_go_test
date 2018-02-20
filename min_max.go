package main

import "fmt"

func  Find_min_max(x []int) (int, int){
	max := x[0] 
	min := x[0] 
        for _, per := range x {
		if max < per {
			max = per
		} else if min > per {
			min = per
		}
	}
	return min, max
}

func main() {
	arr := []int{4,34,12,54,12,43,1,3,12,34,23,12,43,6}
	fmt.Println(Find_min_max(arr))
}