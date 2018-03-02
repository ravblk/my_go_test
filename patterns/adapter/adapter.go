package adapter

type Cater interface {
	Miau() string

}

type Cat struct{
}

func (c *Cat) Miau() string {
	return "Miau"
}
type Doger interface {
	Gav() string

}
type Dog struct{
	*Cat
}


func (d *Dog) Gav() string {
	return d.Miau()
}
