/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start

// Definition for singly-linked list.
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var start, end = 0, len(lists)
	var mid = (start + end) / 2
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 0 {
		return nil
	}
	l1 := mergeKLists(lists[start:mid])
	l2 := mergeKLists(lists[mid:end])
	return mergeTwoLists(l1, l2)
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
