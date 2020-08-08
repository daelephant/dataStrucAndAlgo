/**
 * @Author: yin
 * @Description:palindrome_test
 * @Version: 1.0.0
 * @Time : 2020-08-04 15:36
 */
package linkedlist

import "testing"

func TestIsPalindrome2(t *testing.T) {
	strs := []string{"aba", "heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := NewLinkedList()
		for _, c := range str1 {
			l.InsertToTail(string(c))
		}
		l.Print()
		t.Log(isPalindrome(l))
	}
}

func TestIsPalindrome1(t *testing.T) {
	str1 := "1211"
	l := NewLinkedList()
	for _, c := range str1 {
		l.InsertToTail(string(c))
	}
	l.Print()
	t.Log(isPalindrome1(l.head.next))
}
