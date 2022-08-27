/*
 * @Author: Liang zh.liang.cn@gmail.com
 * @Date: 2022-08-27 15:38:16
 * @LastEditors: Liang zh.liang.cn@gmail.com
 * @LastEditTime: 2022-08-27 16:09:54
 * @Description:
 */
package ds_test

import (
	"fmt"

	"github.com/zh-liang-cn/learning-go/03-datastructure/ds"
)

func ExampleLinkedQueue() {
	q := ds.NewLinkedQueue(3)
	fmt.Println(q.Enque(10)) // 返回 true
	fmt.Println(q.Enque(9))  // 返回 true
	fmt.Println(q.Enque(8))  // 返回 true
	fmt.Println(q.Enque(7))  // 返回 false
	fmt.Println(q.IsFull())  // 返回 true
	fmt.Println(q.Peek())    // 返回 10
	fmt.Println(q.Deque())   // 返回 true
	fmt.Println(q.Peek())    // 返回 9
	fmt.Println(q.Deque())   // 返回 true
	fmt.Println(q.Peek())    // 返回 8
	fmt.Println(q.Deque())   // 返回 true
	fmt.Println(q.Deque())   // 返回 false
	fmt.Println(q.IsEmpty()) // 返回 true
	fmt.Println(q.Peek())    // 返回 <nil>

	// Output:
	// true
	// true
	// true
	// false
	// true
	// 10
	// true
	// 9
	// true
	// 8
	// true
	// false
	// true
	// <nil>
}

func ExampleArrayQueue() {
	q := ds.NewArrayQueue(3)
	fmt.Println(q.Enque(10)) // 返回 true
	fmt.Println(q.Enque(9))  // 返回 true
	fmt.Println(q.Enque(8))  // 返回 true
	fmt.Println(q.Enque(7))  // 返回 false
	fmt.Println(q.IsFull())  // 返回 true
	fmt.Println(q.Peek())    // 返回 10
	fmt.Println(q.Deque())   // 返回 true
	fmt.Println(q.Peek())    // 返回 9
	fmt.Println(q.Deque())   // 返回 true
	fmt.Println(q.Peek())    // 返回 8
	fmt.Println(q.Deque())   // 返回 true
	fmt.Println(q.Deque())   // 返回 false
	fmt.Println(q.IsEmpty()) // 返回 true
	fmt.Println(q.Peek())    // 返回 <nil>

	// Output:
	// true
	// true
	// true
	// false
	// true
	// 10
	// true
	// 9
	// true
	// 8
	// true
	// false
	// true
	// <nil>
}
