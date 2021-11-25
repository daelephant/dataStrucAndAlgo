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
