/**
 * @Author: yin
 * @Description:rev_test
 * @Version: 1.0.0
 * @Time : 2020-08-04 17:50
 */
package linkedlist

import (
	"fmt"
	"testing"
)

func TestRevLinkedList(t *testing.T) {

	str1 := "abb"
	l := NewLinkedList()
	for _, c := range str1 {
		l.InsertToTail(string(c))
	}
	l.Print()
	//t.Log(RevLinkedList(l.head.next))
	//t.Log(ReverseList(l.head.next))
	cur := RevLinkedList(l.head)

	format := ""
	for cur.next != nil {
		format += cur.GetValue().(string)
		cur = cur.next
		format += "---->"
	}
	//format += cur.GetValue().(string)
	fmt.Println(format)

}
