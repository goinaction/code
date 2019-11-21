package main
//main 函数保存在名为main的的包里
//不在 构建工具不会生成可执行的文件
//每个代码文件都属于一个包
//@一个包定义一组编译过的代码，包的名字类似命名空间，可以用来间接访问包内生命的标示符。
//@这个特性把不同的包中的同名标识符区分开


//-sample 外部目录
// -data
//   data.json --数据源
// - matachers
// 	rss.go --搜索rss源的匹配器
// -search 
// 	default.go --搜索默认匹配起
// 	feed.go --用于读取json数据的文件
// 	match.go --用于支持不同匹配器的接口
// 	search.go --执行搜索的主控制逻辑
// main.go		-- 程序的入口

import (//import 导入一段代码
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	//_ 下划线 为了让go对包进行初始化 但是并不使用包里的标识符
	//因为不允许声明不使用 下划线 让编译器接收这类导入，
	//并调用对应的包的init函数
	//目的是调用matcher中的rss.go的init()函数
	"github.com/goinaction/code/chapter2/sample/search"
)//导入search 可以使用search中的run函数

//疑问🤔️ search里的init是声明时间执行的

// init is called prior to main.
//在main之前调用
//标准错误 stderr 标准输出stdout
func init() {
	// Change the device for logging to stdout.
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
//输口文件
func main() {
	// Perform the search for the specified term.
	//使用特定索索项目 search的run函数
	search.Run("president")
}
