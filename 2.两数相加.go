/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
// Definition for singly-linked list.
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var value, add, res int
	var head, node, l *ListNode
	node = &ListNode{
		Val:  0,
		Next: nil,
	}
	head = node
	for l1 != nil && l2 != nil {
		res = l1.Val + l2.Val + add
		value = res % 10
		add = res / 10
		node.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		node = node.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil {
		l = l1
	}
	if l2 != nil {
		l = l2
	}
	for l != nil && add != 0 {
		res = l.Val + add
		value = res % 10
		add = res / 10
		node.Next = &ListNode{
			Val:  value,
			Next: nil,
		}
		node = node.Next
		l = l.Next
	}
	if add != 0 {
		node.Next = &ListNode{
			Val:  add,
			Next: nil,
		}
		node = node.Next
	}
	node.Next = l
	return head.Next
}

// @lc code=end
