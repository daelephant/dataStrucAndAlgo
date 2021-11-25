/**
 * @Author: yin
 * @Description:babble
	空间复杂度为 O(1)，是一个原地排序算法
	相邻的两个元素大小相等的时候，我们不做交换
	冒泡排序是稳定的排序算法
	算法过程动态图，可以看看这个https://visualgo.net/
 * @Version: 1.0.0
 * @Time : 2020-08-11 17:24
*/

package sort

//冒泡排序
/*
冒泡排序是一种简单的排序算法。它重复地走访过要排序的数列，一次比较两个元素，如果它们的顺序错误就把它们交换过来。
走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。

1比较相邻的元素。如果第一个比第二个大，就交换它们两个；
2对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
3针对所有的元素重复以上的步骤，除了最后一个；
重复步骤1~3，直到排序完成。
*/

func BubbleSortV1(nums []int) { //基础版本待优化
	length := len(nums)
	//            arr[n-i, n)已经排好序
	//            通过冒泡在arr[n-i-1]位置放上合适的元素
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
func BubbleSortV1_0(nums []int) { //基础版本1待优化
	length := len(nums)
	//            arr[n-i, n)已经排好序
	//            通过冒泡在arr[n-i-1]位置放上合适的元素
	for i := 0; i < length-1; i++ {
		//            如果内层循环无交换，则已经有序
		isSwapped := false
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
}

func BubbleSortV2(nums []int) { //最优
	//        区别，此时i不再是固定+1
	//        i代表第几轮循环，也代表末尾有i个元素有序
	length := len(nums)
	for i := 0; i < length-1; {
		//            arr[n-i, n)已经排好序
		//            通过冒泡在arr[n-i-1]位置放上合适的元素
		//            内部循环最后交换元素的索引
		lastIndexSwap := 0
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				lastIndexSwap = j + 1
			}
		}
		//            if (lastSwapIndex == 0) break; 和两行后的代码一致
		//            该轮循环后，最后交换元素到数组末尾已经有序
		i = length - lastIndexSwap
	}
}

func BubbleSort(s []int) {
	l := len(s) //数组大小
	if l <= 1 {
		return
	}
	for i := 0; i < l; i++ {
		flag := false //设置标志位，没有交换则代表排好序，主动提前退出
		for j := 0; j < l-i-1; j++ {
			if s[j] > s[j+1] { // 交换
				//s[j], s[j+1] = s[j+1], s[j]
				tmp := s[j]
				s[j] = s[j+1]
				s[j+1] = tmp

				flag = true //有数据交换
			}
		}
		if !flag {
			break //没有数据交换
		}
	}
}

//func BubbleSort(s []int) {
//	l := len(s)
//	if l < 1 {
//		return
//	}
//
//	for i := 0; i < l; i++ {
//		flag := false
//
//		for j := 0; j < l-i-1; j++ { //第二层范围和n和第一层i有关 和本身无关 j<l-i-1
//			if s[j] > s[j+1] {
//				//交换
//				tmp := s[j]
//				s[j] = s[j+1]
//				s[j+1] = tmp
//				flag = true
//			}
//		}
//		if !flag {
//			break
//		}
//	}
//}
