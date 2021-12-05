/**
 * @Author: yin
 * @Description:insertionSort
	将数组中的数据分为两个区间，已排序区间和未排序区间。
	初始已排序区间只有一个元素，就是数组的第一个元素。
	插入算法的核心思想是取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，
	并保证已排序区间数据一直有序。重复这个过程，直到未排序区间中元素为空，算法结束。
 * @Version: 1.0.0
 * @Time : 2020-08-14 19:11
*/

package sort

/*
插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。
它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入

一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：

从第一个元素开始，该元素可以认为已经被排序；
取出下一个元素，在已经排序的元素序列中从后向前扫描；
如果该元素（已排序）大于新元素，将该元素移到下一位置；
重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
将新元素插入到该位置后；
重复步骤2~5。
*/
//循环不变量：
//arr[0,i)已排序  arr[i,n)未排序
//arr[i]放到合适的位置
//有序数组插入O（n）

func InsertV1(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		//将arr[i]插入到合适位置
		t := arr[i]
		j := 0
		for j = i; j-1 >= 0 && arr[j-1] > t; j-- {
			arr[j] = arr[j-1] //元素移位
		}
		arr[j] = t
	}
}

func InsertionSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := a[i] //临时存储数组a的值作为未排序区间的开头
		j := i - 1    //已排序区的结尾
		//查找插入的位置
		for ; j >= 0; j-- { //拿着a[i]从右向左依次比较已排序区
			if a[j] > value {
				a[j+1] = a[j] //数据移动
			} else {
				break
			}
		}
		a[j+1] = value //插入数据
	}
}

func insertTest(nums []int) {
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		var j int
		for j = i; j-1 >= 0 && t < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = t
	}
}
