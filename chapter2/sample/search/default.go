package search
//search 包
// defaultMatcher implements the default matcher.
type defaultMatcher struct{} //使用空结构体声明
//defaultMatcher 类型的结构
//空结构体 不会分配内存 很适合不需要维护状态的类型

// init registers the default matcher with the program.
//被引用时 会被编译器发现，保证在main函数前调用
func init() {//函数将默认匹配器注册到程序里
	var matcher defaultMatcher
	//创建一个defaultMatcher类型值 传递给 search.go register函数
	Register("default", matcher)
}
//将一个matcher值保存到注册匹配的映射中 
//会在main函数之前被完成。 
//init 可以完美的完成这种初始化注册的任务

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
//--->调用方法的例子
//方法声明为使用defaultMatcher类型的值作为接受者
//func (m defaultMatcher) Search(feed *Feed, searchTerm string)
//生命一个指向defaultMatcher类型的指针
//dm := new (defaultMatcher)
//dm.Search(feed, "test")//编译器会揭开dm指针的引用，使用对应的值调用方法
//方法为使用指向defaultMater类型的指针作为接受者
//func (m *defaultMatcher) Search(feed *Feed, searchTerm string)
//var dm defaultMatcher//声明一个 defaultMatcher类型的值
//dm.Search(feed, "test")//编译器会自动生成指针引用dm值，使用指针调用方法
//因大部分方法在调用后都需要维护接受者的状态，所以在一个最佳实践是，将方法的接收者声明为指针
//对于defaultMatcher类型来说使用值作为接收者是因为创建defaultMatcher类型的值不需要分配聂村
//由于defaultMatcher不需要维护状态，所以不需要指针像是接受者

//<----- 👇🤔️没大明白需要测试执行
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