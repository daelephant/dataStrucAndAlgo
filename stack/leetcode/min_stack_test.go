/**
 * @Author: yin
 * @Description:min_stack_test
 * @Version: 1.0.0
 * @Time : 2020-08-11 11:49
 */
package leetcode

import "testing"

var s = Constructor()

func TestMinStack_Push(t *testing.T) {
	s.Push(1)
	s.Push(2)
	s.Push(3)

	t.Log(s)
}

func TestMinStack_Top(t *testing.T) {
	t.Log(s.Top())
}

func TestMinStack_GetMin(t *testing.T) {
	t.Log(s.GetMin())
}

func TestMinStack_Pop(t *testing.T) {
	s.Pop()
	t.Log(s)
}
