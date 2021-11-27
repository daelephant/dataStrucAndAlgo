/**
 * @Author: yin
 * @Description:quick
 * @Version: 1.0.0
 * @Time : 2021/11/26 13:21
 */
package sort

import (
	"math/rand"
)

/*
快排nlog(n) 不稳定
快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，
其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

快速排序使用分治法来把一个串（list）分为两个子串（sub-lists）。具体算法描述如下：

从数列中挑出一个元素，称为 “基准”（pivot）；
重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
*/

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}

	//p := partitionV1(arr, l, r)
	//p := partitionV1Random(arr, l, r)
	p := partitionV2(arr, l, r)
	quickSort(arr, l, p-1)
	quickSort(arr, p+1, r)
}

//单路快排   处理有序数组时非常差
func partitionV1(arr []int, l int, r int) int {
	//        循环不变量，arr[l+1, j] < v, arr[j+1, i] >= v
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

//随机化单路快排 随机选取l    改善处理有序数组，但是处理相同元素的数组性能差
func partitionV1Random(arr []int, l int, r int) int {
	//        生成[l, r]之间的随机索引
	//        处理有序数组时，随机取值，避免每次取到最小值
	//        避免频繁创建Random实例，可从外部传入
	//rand.Seed(time.Now().Unix()) //全局seed 一次就行 使用seed只需要在全局调用一次即可，如果多次调用则有可能取到相同随机数。
	p := l + rand.Intn(r-l+1)
	arr[l], arr[p] = arr[p], arr[l]

	//        循环不变量，arr[l+1, j] < v, arr[j+1, i] >= v
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	//交换 l 和 j的值
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

// 双路快速排序，优化处理相同元素的数组
func partitionV2(arr []int, l int, r int) int {
	p := l + rand.Intn(r-l+1)
	arr[l], arr[p] = arr[p], arr[l]

	//      arr[l+1, i-1] <=v; arr[j+1, r] >=v //        当存在相同元素时，较平均分布在划分点两侧
	i, j := l+1, r
	for { //i和j往中间夹逼
		for i <= j && arr[i] < arr[l] { //保证不越界
			i++
		}
		for j >= i && arr[j] > arr[l] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i] //交换完后继续 夹逼
		i++
		j--
	}
	//交换 l 和 j的值
	arr[l], arr[j] = arr[j], arr[l]
	return j
}
