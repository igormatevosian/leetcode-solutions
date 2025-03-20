class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"

        res = [0] * (len(num1) + len(num2))
        for i in range(len(num1) - 1, -1, -1):
            for j in range(len(num2) - 1, -1, -1):
                product = int(num1[i]) * int(num2[j])
                res[i + j + 1] += product
                res[i + j] += res[i + j + 1] // 10
                res[i + j + 1] %= 10

        while len(res) and not res[0]:
            res.pop(0)

        return "".join(map(str, res))
