package golang

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if index, ok := m[target-v]; ok {
			return []int{k, index}
		}
		m[v] = k
	}
	return []int{-1, -1}
}
