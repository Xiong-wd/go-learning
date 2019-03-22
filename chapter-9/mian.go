// goroutine是一个轻量级的golang协程，
// 任何函数只要加上go 关键字就能送给调度器执行，并且协程和主线程之间的通信是双向的，
// 并且多个协程之间会由调度器选择在合适的时机进行协程的切换

// main方法也是一种协程，所以当协程执行中如果不存在切换协程的机会，可能导致main方法也一直等待导致锁死
// goroutine可能的切换点 1、i/o  select  2、channel 3、 等待锁 4、函数调用 5、runtime.Gosched()

// 协程是一个非抢占式多任务处理
// 协程
package main

import (
	"fmt"
	"time"
)

func main() {
	var arr [10]int
	for i := 0;i< 10 ;i++ {
		go func(i int) {
			for {
				arr[i] ++
				// runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(arr)
}

