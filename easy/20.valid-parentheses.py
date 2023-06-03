#
# @lc app=leetcode id=20 lang=python3
#
# [20] Valid Parentheses
#

# @lc code=start
class Solution:
    def isValid(self, s: str) -> bool:
        # bracketPair maps opening brackets to their closing counterparts
        bracketPair = dict(("()", "[]", "{}"))
        print(bracketPair)

        # stack keeps track of the opening brackets we've seen so far
        stack = []

        for c in s:
            if c in '([{':
                stack.append(c)
            elif len(stack) == 0 or c != bracketPair[stack.pop()]:
                return False
        return len(stack) == 0
# @lc code=end
