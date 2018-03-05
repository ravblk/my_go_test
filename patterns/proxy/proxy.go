package proxy

type Doorer interface {
	Open(string) string
	Close() string	

}

type Door struct {

}

func (d *Door) Open(s string) string {
	return "open ok"
}
func (d *Door) Close() string {
	return "close ok"
}
type Security struct {
	door Doorer
}

func (s *Security) Open(str string) string {
	if str != "key" {
		return "open err"	
	}
	if s.door == nil {
		s.door =  &Door{}
	}

	return s.door.Open("key")
}
func (s *Security) Close() string {
	if s.door == nil {
		s.door =  &Door{}
	}
	return s.door.Close()
}

