/**
 * @Author: yin
 * @Description:binarysearch
 * @Version: 1.0.0
 * @Time : 2021/11/26 15:12
 */

package binary

//二分查找的前提是排好序的 正序 正反写法不同
//    递归实现二分查找

func SearchR(data []int, e int) int {
	return search(data, 0, len(data)-1, e)
}

func search(data []int, l int, r int, target int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if data[mid] == target {
		return mid
	}
	if data[mid] < target {
		return search(data, mid+1, r, target)
	}
	return search(data, l, mid-1, target)
}

//Search   非递归实现二分查找 l r中间位移
func Search(data []int, target int) int {
	l, r := 0, len(data)-1 // 在[l...r]的范围里寻找target
	//不变量 [l,r]
	for l <= r { // 当 l == r时,区间[l...r]依然是有效的
		mid := l + (r-l)/2 //避免整型溢出
		if data[mid] == target {
			return mid
		}
		if data[mid] > target {
			r = mid - 1 // target在[l...mid-1]中;（[mid...r]一定没有target）
		} else {
			l = mid + 1 // target在[mid+1...r]中; ([l...mid]一定没有target)
		}
	}
	return -1
}
