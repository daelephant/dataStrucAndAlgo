/**
 * @Author: yin
 * @Description:middleNode_test
 * @Version: 1.0.0
 * @Time : 2020-08-08 22:26
 */
package leetcode

import (
	"log"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	l := NewLinkedList()
	l.InsertToTail(1)
	l.InsertToTail(2)
	l.InsertToTail(3)
	l.InsertToTail(4)
	l.InsertToTail(5)
	//l.InsertToTail(6)
	l.Print()
	log.Println(l.head.Next)
	node := MiddleNode(l.head.Next)
	node.PrintNode()

}
