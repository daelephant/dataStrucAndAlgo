/**
 * @Author: yin
 * @Description:selection
 * @Version: 1.0.0
 * @Time : 2020-08-15 11:32
 */
package sort

func SelectSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		//查找最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
		//tmp := a[i]
		//a[i] = a[minIndex]
		//a[minIndex] = tmp
	}
}
