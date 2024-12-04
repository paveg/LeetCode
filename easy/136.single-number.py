#
# @lc app=leetcode id=136 lang=python3
#
# [136] Single Number
#

# @lc code=start
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        ans = {}
        res = 0
        for num in nums:
            if num in ans:
                ans[num] += 1
            elif num not in ans:
                ans[num] = 1
        for k, v in ans.items():
            if v == 1:
                res = k
        return res
# @lc code=end
