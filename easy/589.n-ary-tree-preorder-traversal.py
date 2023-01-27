#
# @lc app=leetcode id=589 lang=python3
#
# [589] N-ary Tree Preorder Traversal
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
    def preorder(self, root: 'Node') -> List[int]:
        def o(node: 'Node'):
            if node is None:
                return
            ans.append(node.val)
            for i in range(len(node.children)):
                o(node.children[i])

        ans = []
        o(root)
        return ans


# @lc code=end
