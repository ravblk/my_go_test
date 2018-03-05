package composite

type Assigner interface {
	Add(Assigner)
	Name()string
	TakeTask()[]Assigner
	Print() string
}

type Employee struct {
	name string
}

func (e *Employee) Add(a Assigner)  {

}

func (e *Employee) Name() string  {
	return e.name
}

func (e *Employee) TakeTask() []Assigner  {
	return []Assigner{}
}

func (e *Employee) Print() string  {
	return e.Name()
}

type Team struct {
	name string
	assignees [] Assigner
}


func (t *Team) Add(a... Assigner)  {
	for _, val := range a {
		t.assignees = append(t.assignees, val)
	}
}

func (t *Team) Name() string  {
	return t.name
}

func (t *Team)TakeTask() []Assigner  {
	return t.assignees
}

func (t *Team) Print()string  {
	str := t.Name()
	for _, val := range t.assignees {
		str += val.Name()
	}
	return str
}
