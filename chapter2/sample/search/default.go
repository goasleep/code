package search

// defaultMatcher implements the default matcher.
type defaultMatcher struct{} //空结构体在创建实例的时候，不会分配任何内存

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	Register("default", matcher) //初始化的时候注册默认的匹配器
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {//方法名前带有接受者(m defaultMatcher)，意味着这个方法会和执行接收者的类型绑定在一起
	return nil, nil
}
