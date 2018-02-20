package main

import (
	"fmt"
	"os"
)


func main() {
	file,err := os.Create("test.txt")
	if err != nil {
 	    fmt.Println("Create err ",err)
	    return
	}
	defer file.Close()
	str := "my file"
	buf := []byte(str)
	file.Write(buf)
	stat,err := file.Stat()
	if err != nil {
 	    fmt.Println("Stat err ",err)
	    return
	} 
	fmt.Println("size file",stat.Size())
	readbuf := make([]byte,stat.Size())
	_,err = file.ReadAt(readbuf,0)
	if err != nil {
 	    fmt.Println("Read err ",err)
	    return
	} 
	strout := string(readbuf)
 	fmt.Println(strout)
}