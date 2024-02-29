/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var res = &ListNode{
		Next: head,
	}
	var arr = make([]*ListNode, 0)
	var n = k
	var q = res
	for head != nil {
		arr = append(arr, head)
		head = head.Next
		n--
		if n == 0 {
			for i := k - 1; i >= 0; i-- {
				q.Next = arr[i]
				q = q.Next
			}
			n = k
			arr = make([]*ListNode, 0)
			q.Next = nil
		}
	}
	for i := 0; i < len(arr); i++ {
		q.Next = arr[i]
		q = q.Next
	}
	return res.Next
}

// @lc code=end
