/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var res = &ListNode{
		Next: head,
	}
	var p = res
	for head != nil && head.Next != nil {
		p.Next = head.Next
		head.Next = head.Next.Next
		p.Next.Next = head

		p = head
		head = head.Next
	}
	return res.Next
}

// @lc code=end
