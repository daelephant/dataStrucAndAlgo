/**
 * @Author: yin
 * @Description:middleNode
 * @Version: 1.0.0
 * @Time : 2020-08-08 22:23
 */
package leetcode

func MiddleNode(head *ListNode) *ListNode {
	//[1,2,3,4,5]
	//[1,2,3,4]
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
