/**
 * @Author: yin
 * @Description:209_test
 * @Version: 1.0.0
 * @Time : 2021/11/29 16:24
 */
package array

import (
	"reflect"
	"testing"
)

func TestLeetcode209(t *testing.T) {
	a := []int{2, 3, 1, 2, 4, 3}

	//v := minSubArrayLenV1(7, a)
	//v := minSubArrayLenV2(7, a)
	v := minSubArrayLenT1(7, a)
	expected := 2
	if !reflect.DeepEqual(expected, v) {
		t.Errorf("greeting %+v is not the expected. (%v)", v, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}
