#
# @lc app=leetcode id=217 lang=python3
#
# [217] Contains Duplicate
#

# @lc code=start
class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        m = {}
        ans = False
        for num in nums:
            if num not in m:
                m[num] = 0
            else:
                m[num] += 1
            if m[num] >= 1:
                ans = True
                break
        return ans
        
# @lc code=end

