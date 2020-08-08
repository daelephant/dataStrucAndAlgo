/**
 * @Author: yin
 * @Description:palindrome_test
 * @Version: 1.0.0
 * @Time : 2020-08-07 17:41
 */
package linkedlist

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	l := NewLinkedList()
	l.InsertToHead(1)
	l.InsertToHead(2)
	l.InsertToHead(1)
	l.Print()
	fmt.Println(l.head.value)
	fmt.Println(IsPalindrome(l.head.next))
}
