package main

import (
	"fmt"
	"time"
	)

func my_func_tram(c chan string) {
	for  {
		c <- "tram"	
		}
}

func my_func_pam_pam(c chan string) {
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
	go my_func_tram(c)
	go my_func_pam_pam(c)
        go print(c)

        var input string
        fmt.Scanln(&input)	
	
}