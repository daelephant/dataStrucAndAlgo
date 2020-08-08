/**
 * @Author: yin
 * @Description:base
 * @Version: 1.0.0
 * @Time : 2020-08-08 19:12
 */
package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head   *ListNode
	length uint
}

//func (hea *ListNode) PrintN() {
//	format := ""
//	for hea.Next != nil {
//		format += hea.Val.(string)
//		hea = hea.Next
//		format += "---->"
//	}
//	format += hea.Val.(string)
//	fmt.Println(format)
//}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{v.(int), nil}
}

func (hea *ListNode) GetNext() *ListNode {
	return hea.Next
}

func (hea *ListNode) GetVal() interface{} {
	return hea.Val
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//p节点后插入某个节点
func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}

	newNode := NewListNode(v)
	newNode.Next = p.Next
	p.Next = newNode

	l.length++
	return true
}

//p节点前插入某个节点
func (l LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	//TODO 兼容第一个节点
	if p == nil || p == l.head {
		return false
	}
	cur := l.head.Next
	pre := l.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil {
		return false
	}

	newNode := NewListNode(v)
	pre.Next = newNode
	newNode.Next = cur

	l.length++
	return true
}

//在链表头插入节点
func (l *LinkedList) InsertToHead(v interface{}) bool {
	//l.head为哨兵
	return l.InsertAfter(l.head, v)
}

//在链表尾部插入节点
func (l *LinkedList) InsertToTail(v interface{}) bool {
	cur := l.head
	//移动到链表尾部
	for cur.Next != nil {
		cur = cur.Next
	}

	return l.InsertAfter(cur, v)
}

//通过索引查找节点
func (l *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= l.length {
		return nil
	}

	cur := l.head.Next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

//删除传入的节点
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := l.head.Next
	pre := l.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil {
		return false
	}
	pre.Next = p.Next
	p = nil
	l.length--
	return true
}

//打印
func (l *LinkedList) Print() {
	cur := l.head.Next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetVal())
		cur = cur.Next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

//打印
func (hea *ListNode) PrintNode() {
	cur := hea
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetVal())
		cur = cur.Next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
