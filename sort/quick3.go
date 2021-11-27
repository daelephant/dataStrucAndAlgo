/**
 * @Author: yin
 * @Description:quick3
 * @Version: 1.0.0
 * @Time : 2021/11/26 14:35
 */
package sort

import "math/rand"

func QuickSort3(arr []int) {
	quickSort3(arr, 0, len(arr)-1)
}

// 三路快速排序  优化 把等于v的值放到紧邻左侧
func quickSort3(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := l + rand.Intn(r-l+1)
	arr[l], arr[p] = arr[p], arr[l]

	//循环不变量        arr[l + 1, lt] < v, arr[lt + 1, i - 1] == v, arr[gt, r] > v
	lt, i, gt := l, l+1, r+1
	for i < gt {
		if arr[i] < arr[l] {
			lt++ //腾出<v空间
			arr[lt], arr[i] = arr[i], arr[lt]
			i++ //位移1
		} else if arr[i] > arr[l] {
			gt--                              //腾出>v空间
			arr[gt], arr[i] = arr[i], arr[gt] //原位置数值依然要判断 所以i不用+1
		} else { // ==v 的区间 arr[i] == arr[l]
			i++
		}
	}
	//交换l和lt的值
	arr[l], arr[lt] = arr[lt], arr[l]
	//arr[l, lt - 1] < v, arr[lt, gt - 1] == v, arr[gt, r] > v
	quickSort3(arr, l, lt-1)
	quickSort3(arr, gt, r)
}
