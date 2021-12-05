/**
 * @Author: yin
 * @Description:interview
 * @Version: 1.0.0
 * @Time : 2021/12/2 13:09
 */

package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 8}
	arr2 := []int{3, 9, 10}
	arr := mergeTwoArr(arr1, arr2)
	arr1 = append(arr1, arr2...)
	mergeSort(arr1, 0, len(arr1)-1)
	fmt.Println(arr)
	fmt.Println(arr1)
}

func mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l int, mid int, r int) {
	temp := make([]int, r-l+1)
	copy(temp, arr[l:r+1])

	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = temp[j-l]
			j++
		} else if j > r {
			arr[k] = temp[i-l]
			i++
		} else if temp[i-l] < temp[j-l] {
			arr[k] = temp[i-l]
			i++
		} else {
			arr[k] = temp[j-l]
			j++
		}
	}
}

func mergeTwoArr(arr1 []int, arr2 []int) []int {
	res := make([]int, len(arr1)+len(arr2))
	i, j := 0, 0
	for k := 0; k < len(res); k++ {
		if i > len(arr1)-1 {
			res[k] = arr2[j]
			j++
		} else if j > len(arr2)-1 {
			res[k] = arr1[i]
			i++
		} else if arr1[i] < arr2[j] {
			res[k] = arr1[i]
			i++
		} else {
			res[k] = arr2[j]
			j++
		}
	}
	return res
}
