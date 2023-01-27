# @before-stub-for-debug-begin
from python3problem9 import *
from typing import *
# @before-stub-for-debug-end

#
# @lc app=leetcode id=9 lang=python3
#
# [9] Palindrome Number
#

# @lc code=start
class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False

        listed = list(str(x))
        reverseListed = list(reversed(listed))
        if len(listed) % 2 == 0:
            for i in range(int(len(listed) / 2)):
                if (listed[i] != reverseListed[i]):
                    return False
            return True
        else:
            for i in range(len(listed)):
                if (listed[i] != reverseListed[i]):
                    return False
            return True



# @lc code=end

