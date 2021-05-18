# Given a signed 32-bit integer x, return x with its digits reversed. If reversi
# ng x causes the value to go outside the signed 32-bit integer range [-231, 231 -
#  1], then return 0.
#
#  Assume the environment does not allow you to store 64-bit integers (signed or
#  unsigned).
#
#
#  Example 1:
#  Input: x = 123
# Output: 321
#  Example 2:
#  Input: x = -123
# Output: -321
#  Example 3:
#  Input: x = 120
# Output: 21
#  Example 4:
#  Input: x = 0
# Output: 0
#
#
#  Constraints:
#
#
#  -231 <= x <= 231 - 1
#
#  Related Topics Math
#  👍 4749 👎 7238


# leetcode submit region begin(Prohibit modification and deletion)
class Solution:
    def reverse(self, x: int) -> int:
        ma = 2 ** 31 - 1
        mi = -(2 ** 31)
        chars = list(str(x))
        if chars[0] == '-':
            operator = chars.pop(0)
            res = int(operator + ''.join(list(reversed(chars))))
            if (mi <= res) and (res <= ma):
                return res
            return 0
        res = int(''.join(list(reversed(chars))))
        if (mi <= res) and (res <= ma):
            return res
        return 0
# leetcode submit region end(Prohibit modification and deletion)
