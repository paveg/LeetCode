#
# @lc app=leetcode id=283 lang=python3
#
# [283] Move Zeroes
#

# @lc code=start
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        counter = 0
        for i in range(len(nums)):
            if nums[i] == 0:
                counter += 1
        for i in range(counter):
            nums.remove(0)
            nums.append(0)

        
# @lc code=end

