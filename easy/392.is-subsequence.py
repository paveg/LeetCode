#
# @lc app=leetcode id=392 lang=python3
#
# [392] Is Subsequence
#

# @lc code=start
class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        tidx = 0
        for char in s:
            while tidx < len(t) and t[tidx] != char:
                tidx += 1
            if tidx == len(t):
                return False
            tidx += 1
        return True

# @lc code=end
