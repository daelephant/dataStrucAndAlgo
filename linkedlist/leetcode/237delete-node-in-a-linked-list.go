/**
 * @Author: yin
 * @Description:237delete-node-in-a-linked-list
 * @Version: 1.0.0
 * @Time : 2021/11/29 23:08
 */

package leetcode

//todo step4 链表 链表不仅仅是传真针引线 有时候考虑删除值
//链表不仅仅是传真针引线 有时候考虑删除值
//请编写一个函数，用于 删除单链表中某个特定节点 。在设计函数时需要注意，你无法访问链表的头节点 head ，只能直接访问 要被删除的节点 。
//
//题目数据保证需要删除的节点 不是末尾节点 。

// 237. Delete Node in a Linked List
// https://leetcode.com/problems/delete-node-in-a-linked-list/description/
// 时间复杂度: O(1)
// 空间复杂度: O(1)

func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	if node.Next == nil {
		node = nil
		return
	}
	node.Val = node.Next.Val
	delNode := node.Next
	node.Next = delNode.Next
}

func deleteNodeVx(node *ListNode) {
	// 注意: 这个方法对尾节点不适用。题目中要求了给定的node不是尾节点
	// 我们检查node.next, 如果为null则抛出异常, 确保了node不是尾节点
	if node == nil || node.Next == nil {
		return
	}
	node.Val = node.Next.Val
	delNode := node.Next
	node.Next = delNode.Next
}

//func deleteNode(node *ListNode) {
//	if node == nil || node.Next == nil {
//		return
//	}
//	node.Val = node.Next.Val
//	node.Next = node.Next.Next
//}
