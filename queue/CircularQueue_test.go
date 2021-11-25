/**
 * @Author: yin
 * @Description:CircularQueue_test
 * @Version: 1.0.0
 * @Time : 2020-08-11 16:22
 */
package queue

import "testing"

var cq = NewCircularQueue(7)

func TestCircularQueue_EnQueue(t *testing.T) {
	cq.EnQueue(1)
	cq.EnQueue(2)
	cq.EnQueue(3)
	cq.EnQueue(4)
	cq.EnQueue(5)
	cq.EnQueue(6)

	t.Log(cq)
	cq.DeQueue()
	t.Log(cq)
}
