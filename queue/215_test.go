/**
 * @Author: yin
 * @Description:215_test
 * @Version: 1.0.0
 * @Time : 2021/11/30 22:44
 */
package queue

import (
	"fmt"
	"testing"
)

func TestHeapTopK(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}

	//fmt.Println(findKthLargest(nums, 2))
	//fmt.Println(findKthLargestT(nums, 2))
	//fmt.Println(findKthLargestT1(nums, 2))
	//fmt.Println(findKthLargestT2(nums, 2))
	//fmt.Println(findKthLargestT3(nums, 2))
	//fmt.Println(findKthLargestT4(nums, 2))
	//fmt.Println(findKthLargest1205(nums, 2))
	fmt.Println(findKthLargest12058(nums, 2))

	//binarySearch
	//nums = []int{3, 4, 5, 6, 7, 8, 9, 111}
	//fmt.Println(binarySearch1205(nums, 9))
	//fmt.Println(binarySearch1205U(nums, 9))
}
