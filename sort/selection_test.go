/**
 * @Author: yin
 * @Description:selection_test
 * @Version: 1.0.0
 * @Time : 2020-08-15 14:38
 */
package sort

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	nums := []int{1, 8, 3, 5, 2}

	SelectV1(nums)

	expected := []int{1, 2, 3, 5, 8}
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("greeting %+v is not the expected. (%v)", nums, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}

func TestSelectSort(t *testing.T) {
	s := []int{1, 8, 3, 5, 2}
	SelectSort(s)
	t.Log(s)
}
