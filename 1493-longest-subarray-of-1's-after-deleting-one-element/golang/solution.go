package golang

func longestSubarray(nums []int) int {
	res := 0
	start := 0
	zeros := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeros++
		}
		for zeros > 1 {
			if nums[start] == 0 {
				zeros--
			}
			start++
		}
		res = max(res, i-start)
	}
	return res
}
