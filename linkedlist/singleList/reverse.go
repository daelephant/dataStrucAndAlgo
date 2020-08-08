/**
 * @Author: yin
 * @Description:reverse
 * @Version: 1.0.0
 * @Time : 2020-08-04 19:08
 */
package singleList

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	next *ListNode
	val  int
}

func NewListNode(val int) *ListNode {
	return &ListNode{next: nil, val: val}
}

// 递归的方式
// 递归的关键是退出条件和推导公式： 当node.next==nil的时候，即是退出条件。  推导公式即是：将node的指针反转，并将之前的node节点的next置为空。并返回node的下一个节点。
func ReserveListNodeByRecursion(node *ListNode) *ListNode {
	if node != nil && node.next == nil {
		return node
	}
	newNode := ReserveListNodeByRecursion(node.next)
	node.next.next = node
	node.next = nil

	return newNode
}

// 采用遍历的方式

// 初始化两个变量，pre和cur,  遍历node，将node的next指向前一个及诶单，也就是pre,
// 将pre赋值为本节点，cur为本节点的next节点
func ReserveListNodeByTraverse(node *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode
	for node != nil {
		cur = node.next
		node.next = pre
		pre = node
		node = cur
	}
	return pre
}

func (ld *ListNode) Print() {
	format := ""
	for ld.next != nil {
		format += strconv.Itoa(ld.val)
		ld = ld.next
		format += "---->"
	}
	format += strconv.Itoa(ld.val)
	fmt.Println(format)
}
