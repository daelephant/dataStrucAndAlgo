/**
 * @Author: yin
 * @Description:heap
 * @Version: 1.0.0
 * @Time : 2021/11/30 19:45
 */

package queue

import (
	"container/heap"
	"fmt"
	"testing"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j] //最大堆
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	// Push 和 Pop 使用 pointer receiver 作为参数，
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	//old := *h
	//n := len(old)
	//x := old[n-1]
	//*h = old[:n-1]
	//return x

	//简化写法 弹出slice末尾元素
	del := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return del
}

func TestHeap(t *testing.T) {
	//h := &IntHeap{3, 1, 4, 2}
	//heap.Init(h)
	//h.Push(0)
	h := &IntHeap{}
	heap.Push(h, 2)
	heap.Push(h, 3)
	heap.Push(h, 4)
	heap.Push(h, 2)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))

	}
}

//func TestHeap() {
//	//h := &IntHeap{3, 1, 4}
//	//heap.Init(h)
//	//h.Push(0)
//	//fmt.Println(h.Pop()) //必须使用heap 去操作Pop
//}

//https://blog.csdn.net/zhghost/article/details/108091073?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.no_search_link&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.no_search_link
