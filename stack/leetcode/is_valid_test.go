/**
 * @Author: yin
 * @Description:is_valid_test
 * @Version: 1.0.0
 * @Time : 2020-08-10 20:13
 */
package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	s := "{({)}"
	t.Log(IsValid(s))
}
