#
# @lc app=leetcode id=13 lang=python3
#
# [13] Roman to Integer
#

# @lc code=start
class Solution:
    def romanToInt(self, s: str) -> int:
        result = 0
        mapping: dict[str, int] = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
            'IV': 4,
            'IX': 9,
            'XL': 40,
            'XC': 90,
            'CD': 400,
            'CM': 900,
        }
        listed = list(s)
        caseLength = len(listed)
        index = 0
        if caseLength == 2 and ''.join(listed) in mapping:
            print(''.join(listed))
            return mapping[''.join(listed)]

        while index < caseLength:

            if len(listed) > 1 and (listed[0] + listed[1]) in mapping:
                result += mapping[listed.pop(0) + listed.pop(0)]
                index += 2
                continue
            if len(listed) > 0 and (listed[0]) in mapping:
                result += mapping[listed.pop(0)]
                index += 1

        return result


# @lc code=end

