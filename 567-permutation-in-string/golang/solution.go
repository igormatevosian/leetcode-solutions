package golang

func checkInclusion(s1 string, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		return false
	}
	count1, count2 := [26]int{}, [26]int{}
	for i := 0; i < len1; i++ {
		count1[s1[i]-'a']++
		count2[s2[i]-'a']++
	}
	matches := 0
	for i := 0; i < 26; i++ {
		if count1[i] == count2[i] {
			matches++
		}
	}
	for i := len1; i < len2; i++ {
		if matches == 26 {
			return true
		}
		leftChar := s2[i-len1] - 'a'
		count2[leftChar]--
		if count2[leftChar] == count1[leftChar] {
			matches++
		} else if count2[leftChar]+1 == count1[leftChar] {
			matches--
		}
		rightChar := s2[i] - 'a'
		count2[rightChar]++
		if count2[rightChar] == count1[rightChar] {
			matches++
		} else if count2[rightChar]-1 == count1[rightChar] {
			matches--
		}
	}
	return matches == 26
}
