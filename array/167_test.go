/**
 * @Author: yin
 * @Description:167_test
 * @Version: 1.0.0
 * @Time : 2021/11/29 14:59
 */
package array

import (
	"reflect"
	"testing"
)

// go test -v -run Test
func TestLeetcode167(t *testing.T) {
	a := []int{2, 7, 11, 15}

	v := twoSumV1(a, 9)
	expected := []int{1, 2}
	if !reflect.DeepEqual(expected, v) {
		t.Errorf("greeting %+v is not the expected. (%v)", a, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}
