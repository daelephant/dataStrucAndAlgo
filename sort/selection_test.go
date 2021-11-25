/**
 * @Author: yin
 * @Description:selection_test
 * @Version: 1.0.0
 * @Time : 2020-08-15 14:38
 */
package sort

import "testing"

func TestSelectSort(t *testing.T) {
	s := []int{1, 8, 3, 5, 2}
	SelectSort(s)
	t.Log(s)
}
