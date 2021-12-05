/**
 * @Author: yin
 * @Description:203remove-linked-list-elements
 * @Version: 1.0.0
 * @Time : 2021/11/29 22:15
 */

package leetcode

//todo step2 链表 设立链表的虚拟头结点 Remove Linked List Elements

// 203. Remove Linked List Elements
// https://leetcode.com/problems/remove-linked-list-elements/description/
// 使用虚拟头结点
// 时间复杂度: O(n)
// 空间复杂度: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 使用虚拟头结点
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	// 创建虚拟头结点
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			delNode := cur.Next
			cur.Next = delNode.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

// 不使用虚拟头结点
// 时间复杂度: O(n)
// 空间复杂度: O(1)

func removeElementsV1(head *ListNode, val int) *ListNode {
	// 需要对头结点进行特殊处理
	for head != nil && head.Val == val {
		delNode := head
		head = delNode.Next
	}
	if head == nil {
		return head
	}

	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			delNode := cur.Next
			cur.Next = delNode.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//
func removeElementsT(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{0, head}
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			delNode := cur.Next
			cur.Next = delNode.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
