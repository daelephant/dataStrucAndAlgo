/**
 * @Author: yin
 * @Description:array 插入、删除和按下标随机访问操作 []int
 * @Version: 1.0.0
 * @Time : 2020-08-03 11:28
 */

package array

import (
	"fmt"

	"errors"
)

type Array struct {
	data   []int
	length uint
}

//为数组初始化内存
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity, capacity),
		length: 0,
	}
}

func (a *Array) Len() uint {
	return a.length
}

//判断是否越界
func (a *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(a.data)) { //选取slice的底层数组的cap做最大界定
		return true
	}
	return false
}

//通过索引查找
func (a *Array) Find(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index], nil
}

//指定index插入到数组时间复杂度：min O(1) max:O(n) avg:O(n)
func (a *Array) Insert(index uint, v int) error {
	if a.Len() == uint(cap(a.data)) {
		return errors.New("full array")
	}
	if index != a.length && a.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	//搬运移动index之后的数据
	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

func (a *Array) InsertToTail(v int) error {
	return a.Insert(a.Len(), v)
}

//删除index value
func (a *Array) Delete(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := a.data[index]
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}

//打印
func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
