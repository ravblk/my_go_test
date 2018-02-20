package find

import "testing"



func TestFind_min_max(t *testing.T) {
	type teststruct struct {
		arr []int
        	min int
        	max int
	}

 	testarr := []teststruct {
		{[]int{23,12,32,54,65,121},12,121},
		{[]int{2,12,32,54,65,100},2,100},
	}

	var tmpmin, tmpmax int
	for _,x := range testarr {
		tmpmin, tmpmax = Find_min_max(x.arr)
		if tmpmin != x.min {
			t.Error("for ",x.arr,"min = ",x.min,"got = ", tmpmin)
		}
		if tmpmax != x.max {
			t.Error("for ",x.arr,"max = ",x.max,"got = ", tmpmax)
		}
	}	

}