package golang

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	} else if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	var res []string
	start := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			if start == i-1 {
				res = append(res, strconv.Itoa(nums[start]))
			} else {
				res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[i-1]))
			}
			start = i
		}
	}

	if nums[start] == nums[len(nums)-1] {
		res = append(res, strconv.Itoa(nums[start]))
	} else {
		res = append(res, strconv.Itoa(nums[start])+"->"+strconv.Itoa(nums[len(nums)-1]))
	}
	return res
}
