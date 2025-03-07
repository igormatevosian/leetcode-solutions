package golang

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	var res [][]int
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1])
		}
	}
	return res
}
