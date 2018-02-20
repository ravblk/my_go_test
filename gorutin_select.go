package main

import (
	"fmt"
	"time"
	)



func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "tram"
                        time.Sleep(time.Second * 2)
		}		
	}()
	go func() {
		for {
			c2 <- "pam"
                        time.Sleep(time.Second * 1)
		}		
	}()
        go func() {
		for {
			select {
			case str1 := <- c1:
				fmt.Println(str1)
			case str2 := <- c2:
				fmt.Println(str2)
			}
		}
	}()

        var input string
        fmt.Scanln(&input)	
	
}