class Solution:
    def simplifyPath(self, path: str) -> str:
        stack = []
        for component in path.split("/"):
            if component == "..":
                if stack:
                    stack.pop()
            elif component not in (".", ""):
                stack.append(component)
        return "/" + "/".join(stack)