package main

import "fmt"
import f "my_test/find"


func main() {
	arr := []int{4,34,12,54,0,112,43,12,3,142,34,23,12,43,6}
	fmt.Println(f.Find_min_max(arr))
}