package golang

func generateParenthesis(n int) []string {
	var res []string
	generate(&res, n, 0, 0, "")
	return res
}

func generate(res *[]string, n, left, right int, s string) {
	if left == n && right == n {
		*res = append(*res, s)
	} else if left <= n && right <= n {
		if left < n {
			generate(res, n, left+1, right, s+"(")
		}
		if right < left {
			generate(res, n, left, right+1, s+")")
		}
	}
}
