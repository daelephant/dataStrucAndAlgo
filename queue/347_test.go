/**
 * @Author: yin
 * @Description:347_test
 * @Version: 1.0.0
 * @Time : 2021/11/30 21:35
 */
package queue

import (
	"fmt"
	"testing"
)

func TestTopFreq(t *testing.T) {
	nums := []int{3, 2, 1, 2, 2, 3}
	//m := map[int]int{5: 3, 4: 1, 2: 9}
	topKFrequentV1(nums, 2)
	fmt.Println(nums)
	//quickSort(nums, m)
	//fmt.Println(nums)
	//sort.QuickSort(nums)
	//fmt.Println(nums)

}
