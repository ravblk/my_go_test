package main

import (
	"fmt"
	"container/list"
	"sort"
)


func main() {
	var my list.List
	my_list.PushBack(10)
	my_list.PushBack(12)
	my_list.PushBack(32)
	my_list.PushBack(4)
	my_list.PushFront(0)
	my_list.PushFront(24)
	my_list.PushFront(1)
	mas := make([]int,my.Len())
	j := 0
	for i:= my.Front(); i != nil; i = i.Next() { 	
 		fmt.Println(i.Value.(int))
		mas[j] = i.Value.(int)
		j++
	}

	sort.Ints(mas)	
 	fmt.Println("sort mas = ",mas)
}