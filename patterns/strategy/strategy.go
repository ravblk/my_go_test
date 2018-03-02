package strategy


type Strateger interface {
	Find([]int) int
}



type Min struct {
}

func (m *Min)	Find(k []int) int{
	min := k[0]
	for _,n := range k{
		if min > n{
			min = n
		}
	}
	return min
}

type Max struct {
}

func (m *Max)	Find(k []int) int{
	max := k[0]
	for _,n := range k{
		if max < n{
			max = n
		}
	}
	return max
}

type Context struct {
	srtategy Strateger
}

func (c *Context) Algoritm( s Strateger){
	c.srtategy = s
}

func (c *Context) Find( i []int) int{
	return c.srtategy.Find(i)
}