package decorator

import "fmt"
import 	"testing"

func TestDecorator(t *testing.T) {
	checkMilk := "Cofee with milk"
	costMilk := 60
	m := &MilkCofee{}

	if m.GetDescription() != checkMilk {

		fmt.Println(m.GetDescription(), "err Description")
	}
	if m.GetCost() != costMilk {
		fmt.Println("err Cost")
	}
}				