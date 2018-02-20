package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)


func main() {
	buf,err := ioutil.ReadFile("test.txt")
	if err != nil {
 	    fmt.Println("Create err ",err)
	    return
	}
	hash_crc32 := crc32.NewIEEE()
	hash_crc32.Write(buf)
	fmt.Println(hash_crc32)
}