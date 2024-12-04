#
# @lc app=leetcode id=66 lang=python3
#
# [66] Plus One
#

# @lc code=start
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        length = len(digits)
        n = 0
        for i in range(length):
            n += digits[i] * (10 ** (length - i - 1))
        # @example: 123 -> [1, 2, 3]
        return [int(char) for char in str(n + 1)]
        
# @lc code=end

