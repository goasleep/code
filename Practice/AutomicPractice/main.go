package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64

	wg sync.WaitGroup

	// 定义互斥量
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go increate(1)
	go increate(2)

	wg.Wait()

	fmt.Println("Final counter:", counter)
}

func increate(num int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			value++
			runtime.Gosched()
			counter = value
		}
		mutex.Unlock()

		// atomic.AddInt64(&counter, 1)
		// // 当前goroutine 从线程退出，并放回队列
		// runtime.Gosched()
	}
}
