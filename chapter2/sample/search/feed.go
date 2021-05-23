package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"` // `` 内的部分被称作了标记，用于描述JSON解码的元数据
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {// 最后括号里的是返回值，该函数有两个返回值，第一个是返回值是Feed类型的切片，第二个是error类型的值
	// Open the file.
	file, err := os.Open(dataFile)//调用os的Open，第一个返回值是File类型的指针，第二个是error类型的值，检查文件是否打开成功
	if err != nil {
		return nil, err
	}

	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()//defer 关键字会安排随后的函数调用在函数返回时才执行,被修饰的函数一定会被调用

	// Decode the file into a slice of pointers
	// to Feed values.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// We don't need to check for errors, the caller can do this.
	return feeds, err
}
