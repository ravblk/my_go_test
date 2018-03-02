package strategy

import (
	"testing"
	"fmt"
)

func TestStrategy(t *testing.T) {
	buf := []int{1,2,2,4,5,64,4,3,23,23,32,32,32,232,5,454,43,21,23,5,4}
	
	c := Context{}
	c.Algoritm(&Min{})
	if c.Find(buf) != 1{
		fmt.Println("err min")
	}
	c.Algoritm(&Max{})
	if c.Find(buf) != 454{
		fmt.Println("err max")
	}

}