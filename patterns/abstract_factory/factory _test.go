package abstract_factory

import (
	"testing"
	"fmt"
)

func TestBuilder(t *testing.T) {
        factory := &DoorFactory{}	
	door := factory.CreateDoor(2,3,7)
	master := factory.CreateMaster("lomaster")
	fmt.Println(door.GetWidth(),door.GetHeight(),door.GetType(),master.GetName())			
}


