package search
//search 包
// defaultMatcher implements the default matcher.
type defaultMatcher struct{} //使用空结构体声明
//defaultMatcher 类型的结构
//空结构体 不会分配内存 很适合不需要维护状态的类型

// init registers the default matcher with the program.
func init() {//函数将默认匹配器注册到程序里
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
//Search 实现默认匹配器的行为
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	//search 返回的nil
	return nil, nil 
}
//defaultMatcher类型实现 
//func (m defaultMatcher) Search
//如果声明函数的时候带有接收者，则意味着声明了一个方法。
//这个方法会和指定的接收者的类型绑定在一起
//Search方法 与defaultMatcher 类型值绑定在一起。
//可以使用defaultMatcher 类型的值或者指向这个类型的指针来调用search方法
//无论是接收者类型的值来调用这个方法，还是使用者类型值的指针来调用这个方法
//编译器都会正确的引用对应的值，作为接收者传递给search方法
//例子 绑定使用

//--受限制
//声明为使用指向defaultMatcher类型的指针作为接受者
// func (m *defaultMatcher) Search(feed *Feed, searchTerm string)
// //通过interface类型的值来调用方法
// var dm defaultMatcher
// var matcher Matcher = dm//将值赋值为接口类型
// matcher.Search(feed, "test")//使用值来调用接口方法 

//--通过
// func (m defaultMatcher) Search(feed *Feed, searchTerm strin)
// //通过interface类型的值来调用方法
// var dm defaultMatcher
// var matcher Matcher =&dm//指针赋值给接口类型
// matcher.Search(feed, "test")//指针来调用接口
//