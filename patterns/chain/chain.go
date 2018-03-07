package chain

type Accounter interface {
	SetNext(Accounter) 
	Pay(int) string	

}

type Bank1 struct {
	next Accounter
	many int
}

func (b *Bank1) SetNext(a Accounter)  {
	b.next = a
}
func (b *Bank1) Pay(i int) string{
	if b.many < i  {
		if b.next == nil {
			return "not money on accounts"
		} else {
			return b.next.Pay(i)
		}		
	} else {
		return "bank1 ok"
	}
}

type Bank2 struct {
	next Accounter
	many int
}

func (b *Bank2) SetNext(a Accounter)  {
	b.next = a
}
func (b *Bank2) Pay(i int) string{
	if b.many < i  {
		if b.next == nil {
			return "not money on accounts"
		} else {
			return b.next.Pay(i)
		}		
	} else {
		return "bank2 ok"
	}
}

