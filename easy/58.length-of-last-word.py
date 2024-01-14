#
# @lc app=leetcode id=58 lang=python31
#
# [58] Length of Last Word
#

# @lc code=start
class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        words = s.split(" ")
        words.reverse()
        r = filter(lambda a: a != '', words)
        return len(list(r)[0])


# @lc code=end
