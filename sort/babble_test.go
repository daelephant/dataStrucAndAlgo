/**
 * @Author: yin
 * @Description:babble_test
 * @Version: 1.0.0
 * @Time : 2020-08-11 17:30
 */
package sort

import (
	"reflect"
	"testing"
)

func TestBubbleV(t *testing.T) {
	nums := []int{1, 8, 3, 5, 2}

	//BubbleSort2(nums)
	BubbleSort10(nums)

	expected := []int{1, 2, 3, 5, 8}
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("greeting %+v is not the expected. (%v)", nums, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)
}

func BenchmarkBubbleSortV(b *testing.B) {
	nums := []int{1, 8, 3, 5, 2}
	expected := []int{1, 2, 3, 5, 8}
	for i := 0; i < b.N; i++ {
		//bubbleSort1(nums)
		BubbleSort2(nums)
		if !reflect.DeepEqual(expected, nums) {
			b.Errorf("greeting %+v is not the expected. (%v)", nums, expected)
		}
	}
}

func TestBubbleSortV_(t *testing.T) {
	nums := []int{1, 8, 3, 5, 2}
	bubbleSort1(nums)
	expected := []int{1, 2, 3, 5, 8}
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("greeting %+v is not the expected. (%v)", nums, expected)
	}
	t.Logf("The expected greeting is %v.\n", expected)

	//expected := fmt.Sprintf("Hello, %s!", name)
	//if greeting != expected {
	//	t.Errorf("The actual greeting %q is not the expected. (name=%q)",
	//		greeting, name)
	//}
	//t.Logf("The expected greeting is %q.\n", expected)
}

func BenchmarkBubbleSortV_(b *testing.B) {
	nums := []int{1, 8, 3, 5, 2}
	expected := []int{1, 2, 3, 5, 8}
	for i := 0; i < b.N; i++ {
		bubbleSort1(nums)
		if !reflect.DeepEqual(expected, nums) {
			b.Errorf("greeting %+v is not the expected. (%v)", nums, expected)
		}
	}
}
