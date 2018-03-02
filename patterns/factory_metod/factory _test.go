package factory_method

import (
	"testing"
)

func TestBuilder(t *testing.T) {

				
	assert := []string{"A","B"}
	
	factory := new(ConcreteCreator)

	products := []Producter{
		factory.CreateProduct("A"),	
		factory.CreateProduct("B"),
	}
	for i, product := range products {
		if action := product.Use(); action != assert[i] {
			t.Errorf("Expect action to %s, but %s.\n", assert[i], action)
		}
	}
}


