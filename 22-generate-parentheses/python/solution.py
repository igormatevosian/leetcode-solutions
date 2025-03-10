class Solution:
    def generateParenthesis(self, n: int) -> list[str]:
        def generate(left, right, s):
            if left == right == n:
                result.append(s)
            if left < n:
                generate(left + 1, right, s + '(')
            if right < left <= n:
                generate(left, right + 1, s + ')')

        result = []
        generate(0, 0, '')
        return result
