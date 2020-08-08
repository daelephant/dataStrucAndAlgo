/**
 * @Author: yin
 * @Description:removeNthFromEnd
 * @Version: 1.0.0
 * @Time : 2020-08-08 16:43
 */
package leetcode

import (
	"fmt"
)

//TODO 双层遍历：
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	i, j := 0, 0
	index := 0
	p, ptr := &ListNode{}, &ListNode{}

	p.Next, ptr.Next = head, head

	//寻找尾节点
	for p.Next != nil {
		p = p.Next
		i++
	}
	fmt.Println(i)
	//确定要删除的节点为第几个
	index = i - n + 1

	//将另外一个指针移到index的前一个位置
	for j+1 < index {
		ptr = ptr.Next
		j++
	}

	//删除操作，需要判断要删除的节点是否为第一个节点
	if index == 1 {
		return ptr.Next.Next
	}
	ptr.Next = ptr.Next.Next
	return head
}

//小浩 单层遍历 head遍历到n，开始遍历cur ，并记录cur的pre
func RemoveNthFromEnd1(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result

	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next
}
