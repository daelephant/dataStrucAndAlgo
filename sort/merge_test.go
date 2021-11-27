/**
 * @Author: yin
 * @Description:merge_test
 * @Version: 1.0.0
 * @Time : 2021/11/26 09:17
 */
package sort

import (
	"reflect"
	"testing"
)

//go test -v -run TestMergeSort
func TestMerge(t *testing.T) {
	nums := []int{1, 8, 3, 5, 2}

	MergeSortV1(nums)

	expected := []int{1, 2, 3, 5, 8}
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("greeting %+v is not the expected. (%v)", nums, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}
