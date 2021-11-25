/**
 * @Author: yin
 * @Description:StackBasedOnArray_test
 * @Version: 1.0.0
 * @Time : 2020-08-11 15:35
 */
package stack

import "testing"

var s = NewArrayStack()

func TestArrayStack_Push(t *testing.T) {
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s)
	t.Log(s.Pop())
	t.Log(s.top)
}
