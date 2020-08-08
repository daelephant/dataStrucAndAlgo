/**
 * @Author: yin
 * @Description:removeNthFromEnd_test
 * @Version: 1.0.0
 * @Time : 2020-08-08 19:15
 */
package leetcode

import (
	"log"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	l := NewLinkedList()
	//for i := 0; i < 10; i++ {
	//	l.InsertToHead(i + 1)
	//}
	//l.InsertToHead(3)
	//l.InsertToHead(2)

	l.InsertToTail(1)
	l.InsertToTail(2)
	l.InsertToTail(3)
	l.InsertToTail(4)
	l.InsertToTail(5)
	l.Print()

	RemoveNthFromEnd(l.head.Next, 2)
	log.Println("TestRemoveNthFromEnd after：")
	l.Print()

}

func TestRemoveNthFromEnd1(t *testing.T) {
	l := NewLinkedList()
	//for i := 0; i < 10; i++ {
	//	l.InsertToHead(i + 1)
	//}
	l.InsertToHead(3)
	l.InsertToHead(2)
	l.InsertToHead(1)
	log.Println("TestRemoveNthFromEnd1：")
	l.Print()

	RemoveNthFromEnd1(l.head, 3)
	t.Log("AFTER:")
	l.Print()
}
