//If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


package main

import "fmt"

type Letter struct {
	num map[int]int
}

func (l *Letter) initNum() {
	l.num = make(map[int]int)
        l.num[1] = 3
	l.num[2] = 3
	l.num[3] = 5
	l.num[4] = 4
	l.num[5] = 4
	l.num[6] = 3
	l.num[7] = 5
	l.num[8] = 5
	l.num[9] = 4
	l.num[10] = 3
	l.num[11] = 6
	l.num[12] = 6
	l.num[13] = 8
	l.num[14] = 8
	l.num[15] = 7
	l.num[16] = 7
	l.num[17] = 9
	l.num[18] = 8
	l.num[19] = 8
	l.num[20] = 6
	l.num[30] = 6
	l.num[40] = 5
	l.num[50] = 5
	l.num[60] = 5
	l.num[70] = 7
	l.num[80] = 6
	l.num[90] = 6
}
func (l *Letter)cnt(n int) int {
        str := fmt.Sprintf("%d", n)
	switch len(str) {
	case 0: return 0
	case 1: return l.num[n]
	case 2:
		if el,ok := l.num[n];ok  {
			return el 
		} else {
			k :=  n%10
			return l.num[n-k] + l.num[k]
		}
	case 3:
		k := n/100
		rem := n%100 
		return l.num[k] + 10 + l.cnt(rem)
	case 4:
		k := n/1000
		rem := n%1000
 		return l.num[k] + 8 + l.cnt(rem)
 	default:
	return 0
	}
	return 0
}



func main() {
	var lett Letter 
	var sum int
	lett.initNum() 
	for i := 1;i < 1000 ; i++ {
		sum += lett.cnt(i)		
	}
        fmt.Println(sum)
}