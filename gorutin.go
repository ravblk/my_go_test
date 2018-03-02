package main

import (
	"fmt"
	"time"
	)

func myFunc(n int) {
	for i:= 0;i < 10;i++ {
		fmt.Println("go ",n , " =",i)	
		time.Sleep(time.Duration(100)*time.Millisecond)	
	}
}

func main() {
	go myFunc(0)
	go myFunc(1)
       
        var input string
        fmt.Scanln(&input)	

}