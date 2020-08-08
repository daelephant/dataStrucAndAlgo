/**
 * @Author: yin
 * @Description:rev
 * @Version: 1.0.0
 * @Time : 2020-08-04 17:38
 */
package linkedlist

func RevLinkedList(head *ListNode) *ListNode {
	var new, next *ListNode
	//head.Next, tail, head = tail, head, head.Next
	//return tail
	for head != nil {
		next = head.next
		head.next = new
		new = head
		head = next
	}
	return new
}
