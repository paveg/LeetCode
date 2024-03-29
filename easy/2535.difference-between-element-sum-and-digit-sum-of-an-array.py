#
# @lc app=leetcode id=2535 lang=python3
#
# [2535] Difference Between Element Sum and Digit Sum of an Array
#

# @lc code=start
class Solution:
    def differenceOfSum(self, nums: List[int]) -> int:
        elementSum = sum(nums)
        digitsSum = 0
        for i in range(len(nums)):
            if (len(str(nums[i])) > 1):
                listOfDigits = [int(i) for i in str(nums[i])]
                digitsSum += sum(listOfDigits)
            else:
                digitsSum += nums[i]

        result = elementSum - digitsSum
        return result
# @lc code=end
