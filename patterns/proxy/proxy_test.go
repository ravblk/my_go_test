package proxy

import "fmt"
import 	"testing"

func TestProxy(t *testing.T) {
	check := "open ok"
	d := &Security{}

	if d.Open("key") != check {

		fmt.Println("Err")
	}
	fmt.Println(d.Close())
}				