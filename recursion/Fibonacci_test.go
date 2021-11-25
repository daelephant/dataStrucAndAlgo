/**
 * @Author: yin
 * @Description:Fibonacci_test
 * @Version: 1.0.0
 * @Time : 2020-08-11 16:50
 */
package recursion

import "testing"

func TestFibs_Fibonacci(t *testing.T) {
	fib := NewFibs(10)

	for i := 1; i < 15; i++ {
		fib.Fibonacci(i)
		fib.Print(i)
	}
}
