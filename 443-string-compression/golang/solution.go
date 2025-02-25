package golang

import "strconv"

func compress(chars []byte) int {
	count := 0
	val := chars[0]
	index := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == val {
			count++
		} else {
			chars[index] = val
			index++
			if count > 1 {
				for _, digit := range strconv.Itoa(count) {
					chars[index] = byte(digit)
					index++
				}
			}
			val = chars[i]
			count = 1
		}
	}
	chars[index] = val
	index++
	if count > 1 {
		for _, digit := range strconv.Itoa(count) {
			chars[index] = byte(digit)
			index++
		}
	}
	return index
}
