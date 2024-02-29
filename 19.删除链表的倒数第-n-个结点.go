/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start

// Definition for singly-linked list.
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var l int
	var p = head
	for p != nil {
		p = p.Next
		l++
	}
	count := l - n
	if count == 0 {
		return head.Next
	}
	var q = head
	for i := 1; i < count; i++ {
		q = q.Next
	}
	q.Next = q.Next.Next
	return head
}

// @lc code=end
