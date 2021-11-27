/**
 * @Author: yin
 * @Description:merge
 * @Version: 1.0.0
 * @Time : 2021/11/26 08:46
 */

package sort

/*
归并排序 nlog(n) 稳定  稳定排序只有插入、冒泡和归并。其他的都是o（n+k）级别的桶排序、基数排序，基数排序
归并排序是建立在归并操作上的一种有效的排序算法。
该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并。

把长度为n的输入序列分成两个长度为n/2的子序列；
对这两个子序列分别采用归并排序；
将两个排序好的子序列合并成一个最终的排序序列。
*/

func MergeSortV1(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2

	sort(arr, l, mid)
	sort(arr, mid+1, r)
	//        优化，如果mid小于mid+1，则两个区间已经有序，无需merge 已经 左侧最大值小于右侧最小值
	if arr[mid] > arr[mid+1] {
		merge(arr, l, mid, r)
	}
}

func merge(arr []int, l int, mid int, r int) {
	i := l
	j := mid + 1

	temp := make([]int, r-l+1) //l是偏移量
	copy(temp, arr[l:r+1])
	//每轮循环为arr[k]赋值 在合法区间[l,r]前闭后闭
	for k := l; k <= r; k++ { //            左半部分越界，元素已用完，直接取右半部分
		if i > mid {
			arr[k] = temp[j-l] //l是偏移量 temp取值都要偏移l
			j++
		} else if j > r { //            右半部分越界，元素已用完，直接取左半部分
			arr[k] = temp[i-l]
			i++
		} else if temp[i-l] <= temp[j-l] { //            左部比右部小
			arr[k] = temp[i-l]
			i++
		} else { //            右部比左部小
			arr[k] = temp[j-l]
			j++
		}
	}
}
