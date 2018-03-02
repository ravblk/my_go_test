package factory_method

import "fmt"

type Factorer interface {
	CreateProduct(action string) Producter 
	registerProduct(Producter)            
}


type Producter interface {
	Use() string 
}

type ConcreteCreator struct {
	products []*Producter 
}

func (c *ConcreteCreator) CreateProduct(action string) Producter {
	var product Producter

	switch action {
	case "A":
		product = &ConcreteProductA{action}
	case "B":
		product = &ConcreteProductB{action}
	default:
		fmt.Println("Unknown Action")
	}

	c.registerProduct(product)

	return product
}

func (c *ConcreteCreator) registerProduct(product Producter) {
	c.products = append(c.products, &product)
}

type ConcreteProductA struct {
	action string 
}

func (c *ConcreteProductA) Use() string {
	return c.action
}

type ConcreteProductB struct {
	action string 
}

func (c *ConcreteProductB) Use() string {
	return c.action
}

