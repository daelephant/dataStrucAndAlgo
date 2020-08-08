/**
 * @Author: yin
 * @Description:reverse_test
 * @Version: 1.0.0
 * @Time : 2020-08-04 19:11
 */
package singleList

import (
	"fmt"
	"testing"
)

func TestReverseListNode(t *testing.T) {

	list := NewListNode(0)
	alist := NewListNode(1)
	blist := NewListNode(2)
	clist := NewListNode(3)
	list.next = alist
	alist.next = blist
	blist.next = clist
	list.Print()
	fmt.Println("反转后")
	res := ReserveListNodeByTraverse(list)
	//res:=ReserveListNodeByRecursion(list)
	res.Print()
}
