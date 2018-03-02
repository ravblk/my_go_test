package main

import (
	"fmt"
)

type myStruct struct {
	x int
        y float64
	str string 
}


func (m *myStruct) string() string {
	return fmt.Sprintf("%i %g %s", m.x,m.y,m.str)
}


func main() {
	m := myStruct{1, 2.2, "ir"}
	m1 := myStruct{2, 4.2, "df"}
	fmt.Println(m,m1)

}