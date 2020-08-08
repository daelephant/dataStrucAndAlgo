/**
 * @Author: yin
 * @Description:baseLinkedList 链表的基本操作
 * @Version: 1.0.0
 * @Time : 2020-08-04 11:56
 */
package linkedlist

import (
	"fmt"
)

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func (hea *ListNode) PrintN() {
	format := ""
	for hea.next != nil {
		format += hea.value.(string)
		hea = hea.next
		format += "---->"
	}
	format += hea.value.(string)
	fmt.Println(format)
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (hea *ListNode) GetNext() *ListNode {
	return hea.next
}

func (hea *ListNode) GetValue() interface{} {
	return hea.value
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
	newNode.next = p.next
	p.next = newNode

	l.length++
	return true
}

//p节点前插入某个节点
func (l LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	//TODO 兼容第一个节点
	if p == nil || p == l.head {
		return false
	}
	cur := l.head.next
	pre := l.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}

	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur

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
	for cur.next != nil {
		cur = cur.next
	}

	return l.InsertAfter(cur, v)
}

//通过索引查找节点
func (l *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= l.length {
		return nil
	}

	cur := l.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}
	cur := l.head.next
	pre := l.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = p.next
	p = nil
	l.length--
	return true
}

//打印
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
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
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
