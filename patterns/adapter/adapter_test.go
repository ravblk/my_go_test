package adapter

import "fmt"
import 	"testing"

func TestAdaptor(t *testing.T) {
	dog := &Dog{}
	fmt.Println(dog.Gav())
	if dog.Gav() != "Miau" {
		fmt.Println("err method")
	}
}				