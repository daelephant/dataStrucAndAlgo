/**
 * @Author: yin
 * @Description:19
 * @Version: 1.0.0
 * @Time : 2021/11/29 23:32
 */

package leetcode

// 19. Remove Nth Node From End of List
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
//
// 先记录链表总长度
// 需要对链表进行两次遍历
// 时间复杂度: O(n)
// 空间复杂度: O(1)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	//遍历链表
	length := 0
	for cur := dummyHead.Next; cur != nil; cur = cur.Next {
		length++
	}
	k := length - n
	cur := dummyHead
	//找到待删除指针的前一个node
	for i := 0; i < k; i++ {
		cur = cur.Next
	}
	//删除cur.next  等价与 cur.next = cur.next.next;
	del := cur.Next
	cur.Next = del.Next

	return dummyHead.Next
}

// 使用双指针, 对链表只遍历了一遍
// 时间复杂度: O(n)
// 空间复杂度: O(1)

func removeNthFromEndV1(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}
	p, q := dummyHead, dummyHead
	for i := 0; i < n+1; i++ {
		q = q.Next
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}

	del := p.Next
	p.Next = del.Next

	return dummyHead.Next
}
