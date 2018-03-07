package chain

import "fmt"
import 	"testing"

func TestComposite(t *testing.T) {
	checkStr := "bank2 ok"
	b1 := &Bank1{many:100}
	b2 := &Bank2{many:500}

	b1.SetNext(b2)
	str := b1.Pay(150) 
	if str != checkStr {
		fmt.Println(str)
	}
}				