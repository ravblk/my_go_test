package find


func  Find_min_max(x []int) (int, int){
	max := x[0] 
	min := x[0] 
        for _, per := range x {
		if max < per {
			max = per
		} else if min > per {
			min = per
		}
	}
	return min, max
}

