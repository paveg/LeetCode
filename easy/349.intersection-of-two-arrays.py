#
# @lc app=leetcode id=349 lang=python3
#
# [349] Intersection of Two Arrays
#

# @lc code=start
class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        m = {}
        if len(nums1) > len(nums2):
            for num in nums2:
                if num in nums1:
                    m[num] = True
        else:
            for num in nums1:
                if num in nums2:
                    m[num] = True
        return list(m.keys()) 
# @lc code=end

