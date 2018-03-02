package abstract_factory

type Factorer interface {
	CreateDoor(w int, h int,typeDoor int) AbstractDoor 
	CreateMaster(name string) AbstractMaster            
}


type AbstractDoor interface {
	GetWidth() int
	GetHeight() int
	GetType() int
}

type AbstractMaster interface {
	GetName() string
}


type DoorFactory struct{
}



type Door struct {
	width int
	height int
	doorType int
}

type Master struct {
	name string
}

func (d *DoorFactory) CreateDoor(w int,h int,dt int) *Door {
	return &Door{width:w,height:h,doorType:dt}
}

func (d *DoorFactory) CreateMaster(s string) *Master {
	return &Master{name:s}
}

func (d *Door) GetWidth() int {
	return d.width
}
func (d *Door) GetHeight() int {
	return d.height
}
func (d *Door) GetType() int {
	return d.doorType
}
func (m *Master) GetName() string {
	return m.name
}