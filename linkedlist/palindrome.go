/**
 * @Author: yin
 * @Description:v2_palindrome
 * @Version: 1.0.0
 * @Time : 2020-08-07 14:30
 */
package linkedlist

func IsPalindrome(head *ListNode) bool {
	if head == nil || head.next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	if fast != nil {
		slow = slow.next
	}
	//翻转后半段
	var pre *ListNode
	for slow != nil {
		slow.next, pre, slow = pre, slow, slow.next
	}
	for pre != nil {
		if pre.GetValue() != head.GetValue() {
			return false
		}
		pre = pre.next
		head = head.next
	}
	return true
}
