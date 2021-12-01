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

	fmt.Println(findKthLargest(nums, 2))
}
