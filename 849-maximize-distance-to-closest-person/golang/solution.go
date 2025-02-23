package golang

func maxDistToClosest(seats []int) int {
	res := 0
	start := -1
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if start == -1 {
				res = i
			} else {
				res = max(res, (i-start)/2)
			}
			start = i
		}
	}
	if start != len(seats)-1 {
		res = max(res, len(seats)-1-start)
	}
	return res
}
