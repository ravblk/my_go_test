package command

type Commander interface {
	Execute() string	

}

type TurnOnCommand struct {
	receiver *Receiver
}

func (t *TurnOnCommand) Execute() string {
	return	t.receiver.TurnOn()
}
type TurnOffCommand struct {
	receiver *Receiver
}

func (t *TurnOffCommand) Execute() string {
	return	t.receiver.TurnOff()
}

type  Receiver struct {
}

func (r *Receiver) TurnOn() string  {
	return "On"
}
func (r *Receiver) TurnOff() string  {
	return "Off"
}
type Boss struct {
	commands [] Commander
}

func (b *Boss) StoreCommand(c Commander)  {
	b.commands = append(b.commands,c)
}
func (b *Boss) DelCommand(c Commander)  {
	if b.commands != nil {
		b.commands = b.commands[:len(b.commands)-1]
	}
}
func (b *Boss) Execute()string  {
	var str string
	for _,c := range b.commands  {
		str += c.Execute()
	}
	return str
}