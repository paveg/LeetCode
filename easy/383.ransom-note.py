#
# @lc app=leetcode id=383 lang=python3
#
# [383] Ransom Note
#

# @lc code=start
class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        listed = list(sorted(ransomNote))
        parts = list(sorted(magazine))

        for i in range(len(magazine)):
            if parts[i] == listed[0]:
                listed.pop(0)
                if len(listed) == 0:
                    break
        return len(listed) == 0
# @lc code=end

