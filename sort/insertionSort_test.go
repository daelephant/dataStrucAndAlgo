/**
 * @Author: yin
 * @Description:insertionSort_test
 * @Version: 1.0.0
 * @Time : 2020-08-14 19:20
 */
package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	s := []int{1, 8, 3, 5, 2}
	InsertionSort(s)
	t.Log(s)
}
