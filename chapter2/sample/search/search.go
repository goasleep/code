package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching.
var matchers = make(map[string]Matcher) //包级变量，变量名是以小写开头的，对于小写开头的变量是不公开的。对于大写字母开头是公开的

// Run performs the search logic. 
func Run(searchTerm string) { //func 声明函数，Run函数名，输入的参数为searchTerm，类型为string
	// Retrieve the list of feeds to search through.
	feeds, err := RetrieveFeeds() //调用search包的RetrieveFeeds函数，有两个返回值，一个是Feed类型的切片(动态数组的引用)，第二个为错误值
	if err != nil { // 检测是否有错误
		log.Fatal(err) //log的Fatal函数接收错误值，并将这个错误值在窗口里输出，随后终止程序，自己实现函数，尽量遵守有返回值，也有错误。
	}

	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result) //:= 表示声明一个变量并给其赋值。同样也可以使用var代替，通道和映射和切片一样，都是引用类型

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup //sync的WaitGroup跟踪所有启动的goroutine，是一个计数信号量，用来统计所有的goroutine是不是都完成了工作

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.为每一个feed启动一个goroutine
	for _, feed := range feeds { // range可以迭代数组、字符串、切片、映射和通道。每次迭代会返回元素的索引以及元素值得副本
		// Retrieve a matcher for the search.
		matcher, exists := matchers[feed.Type] // map返回两个值。第一个值为找到得结果，第二个值为查找是否成功，是个布尔值
		if !exists {
			matcher = matchers["default"] // 如果不存在，就会启动默认匹配器。
		}

		// Launch the goroutine to perform the search. 使用关键字启动goroutine
		go func(matcher Matcher, feed *Feed) { // 启用匿名函数
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done() //递减WaitGroup的计数，闭包访问
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()//会导致goroutine阻塞，直到计数器为0

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)//goroutine全部结束了，就可以关闭通道了
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	Display(results)
}

// Register is called to register a matcher for use by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
