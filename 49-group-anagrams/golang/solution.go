package golang

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[string]int)

	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		if index, ok := m[string(runes)]; ok {
			res[index] = append(res[index], str)
		} else {
			res = append(res, []string{str})
			m[string(runes)] = len(res) - 1
		}
	}
	return res
}
