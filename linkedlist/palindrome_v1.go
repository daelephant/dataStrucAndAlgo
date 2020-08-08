/**
 * @Author: yin
 * @Description:palindrome
 * @Version: 1.0.0
 * @Time : 2020-08-04 15:15
 */
package linkedlist

/*
思路：开一个栈存放链表前半段
时间 O(N)
空间 O(N)
*/

func isPalindrome(l *LinkedList) bool {
	lLen := l.length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}

	s := make([]string, 0, lLen/2)
	cur := l.head
	for i := uint(1); i <= lLen; i++ {
		cur = cur.next
		//跳过中间为奇数点的节点
		if lLen%2 != 0 && i == (lLen/2+1) {
			continue
		}
		if i <= lLen/2 {
			s = append(s, cur.GetValue().(string))
		} else {
			if s[lLen-i] != cur.GetValue().(string) {
				return false
			}
		}
	}
	return true
}

/*
思路：
*/

func isPalindrome1(head *ListNode) bool {
	var slow = head
	var fast = head
	var prev *ListNode = nil
	var temp *ListNode = nil

	if head == nil || head.next == nil {
		return true
	}

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		temp = slow.next //temp = cur
		slow.next = prev
		prev = slow
		slow = temp
	} // 快的先跑完,同时反转了一半链表,剪短

	if fast != nil {
		slow = slow.next // 处理余数,跨过中位数
		// prev 增加中 2->1->nil
	}

	var l1 = prev
	var l2 = slow

	for l1 != nil && l2 != nil && l1.GetValue() == l2.GetValue() {
		l1 = l1.next
		l2 = l2.next
	}

	return l1 == nil && l2 == nil

}
