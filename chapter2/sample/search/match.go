package search
//search包
import (
	"log"
)

// Result contains the result of a search.
//保存搜索结果
type Result struct {
	Field   string
	Content string
}

// Matcher defines the behavior required by types that want
// to implement a new search type.
//Matcher 定义了要实现的 新的搜索类型行为
type Matcher interface {//接口类型
	//Search 方法 指向Result 类型值的指针切片 和一个错误值 
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}//Result 在上面已经声明 

//interface 接口类型关键字 
//这个接口类型声明了结构类型或者具名类型需要实现的行为
//接口的行为最终由在这个接口类型中声明的方法来决定
//interface 类型只包含一个方法 这个类型以er结尾
//如果包含多个名字需要和行为关联
//如何实现？ 要实现接口类型里声明的所有方法 =>default.go

// Match is launched as a goroutine for each individual feed to run
//为每个数据源单独启动goroutine来执行这个函数
// searches concurrently. 并发地执行搜索
//Matcher 类型的值作为第一个参数 只有实现了Matcher接口的值或指针才能被接收
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Perform the search against the specified matcher. 对特定匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	// 使用实现了Matcher接口的函数的值或指针进行搜索
	if err != nil {//判断是否为错误
		log.Println(err)
		return
	}

	// Write the results to the channel. 结果写入通道
	for _, result := range searchResults {
		results <- result //结果写入通道
	}//写入结果 --关闭通道--处理结果串在一起
}

// Display writes results to the console window as they
// 从每个goroutine接收结果后在终端打印
// are received by the individual goroutines.
func Display(results chan *Result) {
	// The channel blocks until a result is written to the channel.
	//通道会一直被阻塞，直到有结果写入 一旦通道被关闭 for循环就终止
	// Once the channel is closed the for loop terminates.
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
