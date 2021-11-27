/**
 * @Author: yin
 * @Description:binarysearch_test
 * @Version: 1.0.0
 * @Time : 2021/11/26 15:21
 */
package binary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 20, 66, 77, 80, 1111}

	//r := SearchR(nums, 20)
	r := Search(nums, 20)

	expected := 4
	//if !reflect.DeepEqual(expected, nums) {
	if r != expected {
		t.Errorf("greeting %+v is not the expected. (%v)", r, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}
