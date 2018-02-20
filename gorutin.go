package main

import (
	"fmt"
	"time"
	)

func my_func(n int) {
	for i:= 0;i < 10;i++ {
		fmt.Println("go ",n , " =",i)	
		time.Sleep(time.Duration(100)*time.Millisecond)	
	}
}

func main() {
	go my_func(0)
	go my_func(1)
       
        var input string
        fmt.Scanln(&input)	

}