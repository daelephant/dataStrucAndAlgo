/**
 * @Author: yin
 * @Description:283_test
 * @Version: 1.0.0
 * @Time : 2021/11/27 22:59
 */
package array

import (
	"reflect"
	"testing"
)

// go test -v -run Test
func TestLeetcode283(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}

	//moveZeroes(a)
	moveZeroesV1(a)
	moveZeroesV2(a)
	moveZeroesV3(a)

	expected := []int{1, 3, 12, 0, 0}
	if !reflect.DeepEqual(expected, a) {
		t.Errorf("greeting %+v is not the expected. (%v)", a, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}
