package main

import "fmt"


func main() {
	num := make(map[int]string)
	num[0] = "one"
	num[1] = "two"
	num[2] = "three"
	num[3] = "for"
	num[4] = "five"
	num[5] = "six"
	num[6] = "seven"
	for i := 0;i < len(num); i++ {
		fmt.Println(num[i])
	}
}