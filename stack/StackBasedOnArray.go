/**
 * @Author: yin
 * @Description:StackBasedOnArray
 * @Version: 1.0.0
 * @Time : 2020-08-11 15:01
 */
package stack

import "fmt"

type ArrayStack struct {
	data []interface{} //数据
	top  int           //栈顶指针(索引)
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (a *ArrayStack) IsEmpty() bool {
	if a.top < 0 {
		return true
	}
	return false
}

func (a *ArrayStack) Push(v interface{}) {
	if a.top < 0 {
		a.top = 0
	} else {
		a.top += 1
	}

	if a.top > len(a.data)-1 {
		a.data = append(a.data, v)
	} else {
		a.data[a.top] = v
	}
}

func (a *ArrayStack) Pop() interface{} {
	if a.IsEmpty() {
		return nil
	}
	v := a.data[a.top]
	a.top -= 1
	return v
}

func (a ArrayStack) Top() interface{} {
	if a.IsEmpty() {
		return nil
	}
	return a.data[a.top]
}

func (a *ArrayStack) Flush() {
	a.top = -1
}

func (a *ArrayStack) Print() {
	if a.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := a.top; i > 0; i-- {
			fmt.Println(a.data[i])
		}
	}
}
