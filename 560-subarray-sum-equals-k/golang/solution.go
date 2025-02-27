package golang

func subarraySum(nums []int, k int) int {
	prefixSums := make(map[int]int)
	prefixSums[0] = 1
	prefixSum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		prefixSum += nums[i]
		if el, ok := prefixSums[prefixSum-k]; ok {
			res += el
		}
		prefixSums[prefixSum]++
	}
	return res
}
