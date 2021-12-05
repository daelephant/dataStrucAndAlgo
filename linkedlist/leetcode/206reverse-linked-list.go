/**
 * @Author: yin
 * @Description:206reverse-linked-list
 * @Version: 1.0.0
 * @Time : 2021/11/29 21:27
 */

package leetcode

//todo step1 链表 Reverse Linked List

// 206. Reverse Linked List
// https://leetcode.com/problems/reverse-linked-list/description/
// 时间复杂度: O(n)
// 空间复杂度: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	//链表设置节点为NULL 直接声明即可
	var pre *ListNode //注意不能使用短变量 会默认初始化0值
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 递归的方式反转链表
// 时间复杂度: O(n)
// 空间复杂度: O(n) - 注意，递归是占用空间的，占用空间的大小和递归深度成正比：）

func reverseListV1(head *ListNode) *ListNode {
	// 递归终止条件
	if head == nil || head.Next == nil {
		return head
	}
	rHead := reverseListV1(head.Next)
	head.Next.Next = head
	head.Next = nil

	return rHead
}

func reverseListT(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next //定义next

		cur.Next = pre //转换指针
		//前移
		pre = cur
		cur = next
	}
	return pre
}

//递归
func reverseListT1(head *ListNode) *ListNode {

	//终止
	if head == nil || head.Next == nil {
		return head
	}

	rHead := reverseListT1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rHead
}
