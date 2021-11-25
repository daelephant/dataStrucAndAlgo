/**
 * @Author: yin
 * @Description:QueueBasedOnArray_test
 * @Version: 1.0.0
 * @Time : 2020-08-11 15:53
 */
package queue

import "testing"

var q = NewArrayQueue(5)

func TestArrayQueue_EnQueue(t *testing.T) {
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)

	t.Log(q)

	q.Dequeue()
	t.Log(q)

	t.Log(q.String())
}
