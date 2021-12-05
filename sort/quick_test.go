/**
 * @Author: yin
 * @Description:quick_test
 * @Version: 1.0.0
 * @Time : 2021/11/26 13:51
 */
package sort

import (
	"reflect"
	"testing"
)

//go test -v -run TestQuick
func TestQuick(t *testing.T) {
	nums := []int{1, 8, 3, 5, 2}

	//QuickSort(nums)
	//quickTest(nums)
	quick3Test(nums) //12-2
	//QuickSort3(nums)
	//array.Sort3(nums)

	expected := []int{1, 2, 3, 5, 8}
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("greeting %+v is not the expected. (%v)", nums, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}
