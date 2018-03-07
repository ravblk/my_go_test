package command


import	"testing"

func TestCommand(t *testing.T) {

	chain := "OnOff"
	
	b := &Boss{}
	r := &Receiver{}
	
	b.StoreCommand(&TurnOnCommand{receiver:r})
	b.StoreCommand(&TurnOffCommand{receiver:r})

	
	res := b.Execute()
	
	if res != chain {
		t.Errorf("Err", res)
	}
}