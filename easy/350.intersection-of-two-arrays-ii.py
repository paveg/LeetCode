#
# @lc app=leetcode id=350 lang=python3
#
# [350] Intersection of Two Arrays II
#

# @lc code=start
class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        # MEMO: Another solution using `collections.Counter`
        # c = collections.Counter(nums1) & collections.Counter(nums2)
        # ans = []
        # for k,v in c.items():
        #     ans.extend([k] * v)
        # return ans
        ans = []
        if len(nums1) > len(nums2):
            for num in nums2:
                if num in nums1:
                    ans.append(num)
                    nums1.remove(num)
        else:
            for num in nums1:
                if num in nums2:
                    ans.append(num)
                    nums2.remove(num)
        return ans     

# @lc code=end

