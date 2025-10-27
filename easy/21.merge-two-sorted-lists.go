package code

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := &ListNode{}
	head := curr

	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				curr.Next = list1
				list1 = list1.Next
			} else {
				curr.Next = list2
				list2 = list2.Next
			}
			curr = curr.Next
		} else if list1 != nil {
			curr.Next = list1
			list1 = list1.Next
			curr = curr.Next
		} else if list2 != nil {
			curr.Next = list2
			list2 = list2.Next
			curr = curr.Next
		}
	}

	return head.Next
}

// @lc code=end
