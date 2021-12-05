/**
 * @Author: yin
 * @Description:24swap-nodes-in-pairs
 * @Version: 1.0.0
 * @Time : 2021/11/29 22:56
 */

package leetcode

//todo step3 链表 复杂的穿针引线 Swap Nodes in Pairs 24. 两两交换链表中的节点

// 24. Swap Nodes in Pairs
// https://leetcode.com/problems/swap-nodes-in-pairs/description/
// 时间复杂度: O(n)
// 空间复杂度: O(1)

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	p := dummyHead
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := node1.Next
		next := node2.Next

		node2.Next = node1
		node1.Next = next
		p.Next = node2

		p = node1 //链表中的位移 是直接赋值
	}
	return dummyHead.Next
}

func swapPairsT(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	p := dummyHead
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := node1.Next
		next := node2.Next //赋值注意

		node2.Next = node1
		node1.Next = next
		p.Next = node2

		p = node1
	}
	return p.Next
}
