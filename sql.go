package main

import (
	"fmt"
	"os"
	"mysql"
)

 
func main(){
	user := "root"
	pass := "root"
	dbname := "my_family"
	proto := "tcp"
	addr := "127.0.0.1:3306"

	db := mysql.New(proto, "", addr, user, pass, dbname)

	fmt.Println("Connect to %s:%s... ", proto, addr)


}