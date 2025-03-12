package golang

func isValid(s string) bool {
	var stack []rune
	bracketPairs := map[rune]rune{}
	bracketPairs['('] = ')'
	bracketPairs['['] = ']'
	bracketPairs['{'] = '}'
	for _, bracket := range s {
		if closed, ok := bracketPairs[bracket]; ok {
			stack = append(stack, closed)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != bracket {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
