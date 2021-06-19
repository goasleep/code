package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 分配1个逻辑处理器
	runtime.GOMAXPROCS(2)

	//wg 使用WaitGroup等待goroutine完成，类似于信号量
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutine")

	go func() {
		defer wg.Done() // 保证每个工作完成都调用wg.Done方法

		for count := 0; count < 2; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}() // 后加括号表示立即调用

	go func() {
		defer wg.Done()

		for count := 0; count < 2; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Printf("wait for finish\n")

	wg.Wait()

	fmt.Printf("\n terminal")
}
