/**
 * @Author: yin
 * @Description:mergeTwoLists
 * 小浩 哨兵节点
 * @Version: 1.0.0
 * @Time : 2020-08-08 15:38
 */
package leetcode

//func mergeTwoLists(l1, l2 *ListNode) *ListNode {
//	prehead := &ListNode{}
//	result := prehead
//
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			prehead.Next = l1
//			l1 = l1.Next
//		} else {
//			prehead.Next = l2
//			l2 = l2.Next
//		}
//		prehead = prehead.Next
//	}
//
//	if l1 != nil {
//		prehead.Next = l1
//	}
//	if l2 != nil {
//		prehead.Next = l2
//	}
//	return result.Next
//}
//
////递归
//func mergeTwoLists1(l1, l2 *ListNode) *ListNode {
//	if l1 == nil {
//		return l2
//	}
//	if l2 == nil {
//		return l1
//	}
//
//	if l1.Val < l2.Val {
//		l1.Next = mergeTwoLists1(l1.Next, l2)
//		return l1
//	} else {
//		l2.Next = mergeTwoLists1(l1, l2.Next)
//		return l2
//	}
//}
