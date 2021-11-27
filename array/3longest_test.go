/**
 * @Author: yin
 * @Description:3longest_test
 * @Version: 1.0.0
 * @Time : 2021/11/26 22:58
 */
package array

import (
	"reflect"
	"testing"
)

// go test -v -run Test
func TestLeetcode3(t *testing.T) {
	s := "a"

	get := LengthOfLongestSubstring(s)

	expected := 2
	if !reflect.DeepEqual(expected, get) {
		t.Errorf("greeting %+v is not the expected. (%v)", get, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}
