#
# @lc app=leetcode id=125 lang=python3
#
# [125] Valid Palindrome
#

# @lc code=start
class Solution:
    def isPalindrome(self, s: str) -> bool:
        processed = s.lower()
        processed = ''.join(filter(lambda a: a.isalnum(), processed))
        return processed == processed[::-1]

# @lc code=end
