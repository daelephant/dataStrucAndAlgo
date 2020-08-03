/**
 * @Author: yin
 * @Description:array_test
 * @Version: 1.0.0
 * @Time : 2020-08-03 14:10
 */
package array

import (
	"testing"
)

func TestArray_Insert(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity-2; i++ {
		err := arr.Insert(uint(i), i+1)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	arr.Print()
	arr.Insert(uint(6), 999)
	arr.Print()

	arr.InsertToTail(666)
	arr.Print()
}

func TestArray_Delete(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))
		if nil != err {
			t.Fatal(err)
		}
		arr.Print()
	}
}

func TestArray_Find(t *testing.T) {
	capacity := 10
	arr := NewArray(uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	t.Log(arr.Find(0))
	t.Log(arr.Find(9))
	t.Log(arr.Find(11))
}
