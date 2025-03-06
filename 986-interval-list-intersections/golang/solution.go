package golang

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	first, second := 0, 0
	for first < len(firstList) && second < len(secondList) {
		if firstList[first][1] < secondList[second][0] {
			first++
		} else if secondList[second][1] < firstList[first][0] {
			second++
		} else {
			start := max(firstList[first][0], secondList[second][0])
			end := min(firstList[first][1], secondList[second][1])
			res = append(res, []int{start, end})
			if end < firstList[first][1] {
				second++
			} else {
				first++
			}
		}
	}
	return res
}
