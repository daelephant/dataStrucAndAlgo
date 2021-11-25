/**
 * @Author: yin
 * @Description:min_stack
 * @Version: 1.0.0
 * @Time : 2020-08-11 11:36
 */
package leetcode

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

//初始化
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (m *MinStack) Push(x int) {
	m.stack = append(m.stack, x)
	top := m.minStack[len(m.minStack)-1]
	m.minStack = append(m.minStack, min(x, top))
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
	m.minStack = m.minStack[:len(m.minStack)-1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
