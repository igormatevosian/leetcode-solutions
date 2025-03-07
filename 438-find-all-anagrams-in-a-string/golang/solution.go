package golang

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	res := make([]int, 0)
	mP := make([]int, 26)
	mS := make([]int, 26)
	for i := 0; i < len(p); i++ {
		mP[p[i]-'a']++
	}
	start := 0
	for i := 0; i < len(s); i++ {
		mS[s[i]-'a']++
		if i-start == len(p)-1 {
			if isEqual(mS, mP) {
				res = append(res, start)
			}
			mS[s[start]-'a']--
			start++
		}
	}
	return res
}

func isEqual(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
