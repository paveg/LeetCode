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
        for i in range(len(nums)):
            if nums[i] in ans:
                ans[nums[i]] += 1
            elif nums[i] not in ans:
                ans[nums[i]] = 1
        for k, v in ans.items():
            if v == 1:
                res = k
        return res
# @lc code=end
