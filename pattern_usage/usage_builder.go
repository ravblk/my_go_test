package main

import (
	"fmt"
)

type House struct {
	room int
	window int
	floor int
	color string
}

type Builder interface {
	makeColor(s string) 
	cntRoom(i int) 
	cntWindow(i int) 
	cntFloor(i int) 
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.makeColor("white")
	d.builder.cntRoom(1)
	d.builder.cntWindow(0)
	d.builder.cntFloor(1)
}

type ConcreteBuilder struct {
	house *House
}


func (h *ConcreteBuilder) cntRoom(i int)  {
	h.house.room = i
}

func (h *ConcreteBuilder) cntWindow(i int)  {
	h.house.window = i

}
func (h *ConcreteBuilder) cntFloor(i int)  {
	h.house.floor = i

}

func (h *ConcreteBuilder) makeColor(s string)  {
	h.house.color = s

}


func (h *House) Show() {
fmt.Println("House: cnt room =", h.room,
		" cnt window =", h.window,
		" cnt floor =", h.floor,
		"color",h.color)
}


func main() {
	

	house := &House{}
	
	dir := Director{&ConcreteBuilder{house}}
	dir.builder.makeColor("white")
	dir.builder.cntRoom(10)
	dir.builder.cntWindow(20)
	dir.builder.cntFloor(1)

	house.Show()
}