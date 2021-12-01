/**
 * @Author: yin
 * @Description:24swap-nodes-in-pairs
 * @Version: 1.0.0
 * @Time : 2021/11/29 22:56
 */

package leetcode

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

		p = node1
	}
	return dummyHead.Next
}
