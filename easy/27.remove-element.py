#
# @lc app=leetcode id=27 lang=python3
#
# [27] Remove Element
#

# @lc code=start
class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        j = 0
        for k in nums:
            if k != val:
                nums[j] = k
                j += 1
        return j

# @lc code=end
