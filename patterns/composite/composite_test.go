package composite

import "fmt"
import 	"testing"

func TestComposite(t *testing.T) {
	checkStr := "t1e1e2e3"
	e1 := &Employee{name:"e1"}
	e2 := &Employee{name:"e2"}
	e3 := &Employee{name:"e3"}
	t1 := &Team{name:"t1"}
	t1.Add(e1,e2,e3)
	if t1.Print() != checkStr {
		fmt.Println(t1.Print())
	}
}				