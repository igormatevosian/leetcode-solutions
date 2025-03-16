class Solution:
    def evalRPN(self, tokens: list[str]) -> int:
        stack = []

        for token in tokens:
            if len(token) > 1 or token.isnumeric():
                stack.append(int(token))
            else:
                second = stack.pop()
                first = stack.pop()
                if token == "+":
                    stack.append(first + second)
                elif token == "*":
                    stack.append(first * second)
                elif token == "-":
                    stack.append(first - second)
                else:
                    stack.append(int(first / second))

        return stack.pop()
