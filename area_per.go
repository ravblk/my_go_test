package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) per() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) dist1() float64 {
	a := r.x1 - r.x1
	b := r.y2 - r.y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) dist2() float64 {
	a := r.x2 - r.x1
	b := r.y1 - r.y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	return r.dist1() * r.dist2()
}

func (r *Rectangle) per() float64 {
	return (r.x2-r.x1)*2 + (r.y2-r.y1)*2
}

type Shape interface {
	area() float64
	per() float64
}

func totalAreaPer(shapes ...Shape) (float64, float64) {
	var area, per float64
	for _, s := range shapes {
		area += s.area()
		per += s.per()
	}
	return area, per
}

func main() {
	c := Circle{0, 0, 3}
	r := Rectangle{0, 0, 2, 7}
	c1 := Circle{0, 0, 6}
	r1 := Rectangle{0, 0, 10, 5}

	totalArea, totalPer := totalAreaPer(&c, &r, &c1, &r1)
	fmt.Println(totalArea, totalPer)

}