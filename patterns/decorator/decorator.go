package decorator

type Coffer interface {
	GetCost() int
	GetDescription() string  

}

type Cofee struct {

}


func (c *Cofee) GetCost() int {
	return 50
}

func (c *Cofee) GetDescription() string {
	return "Cofee"
}

type MilkCofee struct {
	cofee *Cofee
}


func (m *MilkCofee) GetCost() int {
	return m.cofee.GetCost() + 10
}

func (m *MilkCofee) GetDescription() string {
	return m.cofee.GetDescription() + " with milk"
}


