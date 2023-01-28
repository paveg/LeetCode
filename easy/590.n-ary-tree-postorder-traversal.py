#
# @lc app=leetcode id=590 lang=python3
#
# [590] N-ary Tree Postorder Traversal
#

# @lc code=start
"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""


class Solution:
    def postorder(self, root: 'Node') -> List[int]:
        def ary_postorder(node: 'Node'):
            if node is None:
                return
            for i in range(len(node.children)):
                ary_postorder(node.children[i])
            ans.append(node.val)

        ans = []
        ary_postorder(root)
        return ans
# @lc code=end
