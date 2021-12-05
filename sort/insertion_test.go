/**
 * @Author: yin
 * @Description:insertionSort_test
 * @Version: 1.0.0
 * @Time : 2020-08-14 19:20
 */
package sort

import (
	"reflect"
	"testing"
)

// go test -v -run TestInsert
func TestInsert(t *testing.T) {
	nums := []int{1, 8, 3, 5, 2}

	//InsertV1(nums)
	insertTest(nums)

	expected := []int{1, 2, 3, 5, 8}
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("greeting %+v is not the expected. (%v)", nums, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}

func TestInsertionSort(t *testing.T) {
	s := []int{1, 8, 3, 5, 2}
	InsertionSort(s)
	t.Log(s)
}
