/**
 * @Author: yin
 * @Description:has_cycle
 * @Version: 1.0.0
 * @Time : 2020-08-08 14:44
 */
package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//字典法
//func hasCycle(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//		return false
//	}
//	//用map暂存
//	tmp := map[*ListNode]interface{}{}
//	for head.Next != nil {
//		if _, ok := tmp[head]; ok {
//			return true
//		}
//		tmp[head] = struct{}{}
//		head = head.Next
//	}
//	return false
//
//}

//快慢指针

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycleT(head *ListNode) bool {
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s.Val == f.Val {
			return true
		}
	}
	return false
}
