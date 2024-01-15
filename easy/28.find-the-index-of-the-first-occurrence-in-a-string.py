#
# @lc app=leetcode id=28 lang=python3
#
# [28] Find the Index of the First Occurrence in a String
#

# @lc code=start
class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        needle_len = len(needle)
        haystack_len = len(haystack)
        result = -1
        if needle_len == 0:
            result = 0
        if needle_len == haystack_len:
            if needle == haystack:
                result = 0
            else:
                result = -1
        for i in range(haystack_len - needle_len + 1):
            if haystack[i:i+needle_len] == needle:
                result = i
                break

        return result

# @lc code=end
