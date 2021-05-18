#
# @lc app=leetcode id=7 lang=python3
#
# [7] Reverse Integer
#

# @lc code=start
class Solution:
    def reverseUtil(self, chars: List[str]) -> int:
        if chars[0] == '-':
            op = chars.pop(0)
            return int(op + ''.join(list(reversed(chars))))
        return int(''.join(list(reversed(chars))))

    def reverse(self, x: int) -> int:
        chars = list(str(x))
        result = self.reverseUtil(chars)
        if (2 ** 31) > result and result >= -(2 ** 31):
            return result
        else:
            return 0
# @lc code=end

