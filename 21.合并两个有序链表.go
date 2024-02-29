/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start

// Definition for singly-linked list.
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var res = &ListNode{}
	var head = res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			res.Next = list1
			list1 = list1.Next
		} else {
			res.Next = list2
			list2 = list2.Next
		}
		res = res.Next
	}
	if list1 != nil {
		res.Next = list1
	}
	if list2 != nil {
		res.Next = list2
	}

	return head.Next
}

// @lc code=end
