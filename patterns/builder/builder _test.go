package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {

				
	house := &House{}
	
	director := Director{&ConcreteBuilder{house}}
	director.Construct()

	house.Show()
	

}


