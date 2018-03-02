package main

import (
	"fmt"
	"time"
	)

func tram(c chan string) {
	for  {
		c <- "tram"	
		}
}

func pam(c chan string) {
	for  {
		c <- "pam pam"	
		}
}

func print(c chan string){
        for{
	msg := <- c
        fmt.Println(msg)	
	time.Sleep(1*time.Second)
	}
}	


func main() {
	var c chan string = make(chan string)
	go tram(c)
	go pam(c)
        go print(c)

        var input string
        fmt.Scanln(&input)	
	
}