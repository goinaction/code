package search
//package +包名字
//search 文件夹下都以search作为包名
import (
	"log"//导入log sync包 stdout stderr
	"sync"//标准库  同步goroutine功能
)//编译器查找包时会到Goroot Gopath环境变量的引用位置去查找

// A map of registered matchers for searching.
//注册搜索的匹配器的映射
var matchers = make(map[string]Matcher)
// matchers没有在任何作用域内，是包级变量 
//关键var声明  小写开头,非公开 var声明是初始化为零0的
//make 构造map类型 map是一个引用类型要用make构造
//map 默认的零值是nil 不构造直接用报错
//类型是map 映射以string类型值作为键
// Matacher类型值作为映射后的值
//Matcher 类型代表文件matcher.go中的声明 type Matcher{}

//@所有变量初始化为其零值。数值是0，字符串空字符串，布尔类型，false，指针，nil
//@引用类型引用的底层数据结构被初始化为零值
//声明为其零值的引用类型返回nil

// @标识符（变量）要么从包里公开，要么不从包里公开
// @大写标识符公开，小写字母开头的不公开，不能被其他包中代码直接访问
// 可以使用一个反问未公开类型的值的函数进行问访问 非公开的标识符

// Run performs the search logic.
//func 声明函数 
//func 函数名（参数，返回值）
func Run(searchTerm string) {//一个string类型参数
	// Retrieve the list of feeds to search through.
	//获取资源数据列表RetrieveFeeds（）方法在feed中
	feeds, err := RetrieveFeeds()
	//：= 变量声明运算符 声明并赋值
	//search.RetrieveFeeds([]*fedds,error)
	//如果错误调用Log.fatal函数
	if err != nil {
		log.Fatal(err)
	}
	//函数返回错误和另一个值，如果返回错误则不要使用另一个值

	// Create an unbuffered channel to receive match results to display.
	//make 一个chan无缓冲区通道 接收匹配结果
	results := make(chan *Result)
	//如果声明初始化为0的使用var 关键字
	//如果非零的或函数函数返回的使用简化声明运算符 :=

	//channel map slice 都是引用类型

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup//处理所有的数据源
	//sync.WaitGroup跟踪所有启动的goroutine
	//WaitGroup是一个技术信号量，统计goroutine是否完成了
	// waitGroup 类型变量
	// Set the number of goroutines we need to wait for while
	//设置需要等待处理 每个数据源的goroutine数量
	// they process the individual feeds.
	waitGroup.Add(len(feeds))
	//设置waitGroup 值为要启动的goroutine的数量
	//goroutine完成后会 递减waitGroup的变量计数值 为0时完成

	// Launch a goroutine for each feed to find the results.
	//为每个数据源启动一个goroutine来查找结果
	//feed matcher 会随着循环迭代而改变
	for _, feed := range feeds {
		//关键字 for range 对feeds切片迭代
		//range用于迭代数组，字符串，切片，映射，通道
		//for range 返回两个第一个是索引位置，第二个是值的副本
		//_ 下划线 用来占位

		// Retrieve a matcher for the search.
		//获取一个匹配器来查找 根据数据源类型查找一个匹配器值
		//给下面的goroutine使用
		matcher, exists := matchers[feed.Type]
		//map检查是否含有数据的类型
		//查找map里的键时，要么赋值给一个变量，
		//要么精确查找赋值给两个变量
		//2个变量时 第一个值就是查找结果，第二时布尔标志位，
		//不存在返回 零值 ，存在返回键所对应值的副本
		if !exists {
			matcher = matchers["default"]
		}//不存在使用默认匹配器

		// Launch the goroutine to perform the search.
		//启动一个goroutine 来执行搜索
		//匿名函数 👇 两个变量参数 matcher feed指针
		//
		go func(matcher Matcher, feed *Feed) {
			//go关键字 启动goroutine 做并发调度
			//Match.go 中的Match函数
			//match函数参数matcher 指向Feed的指针，搜索项 通道
			//疑问🤔️searchTerm是怎么使用的
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()//闭包，函数可以访问没有作为参数传入的变量
			//waitgroup的值没有作为参数传入
			//通过闭包访问的searchTerm, results
			// 是访问外层函数作用域中声明这些变量的本身
		//因为matcher feed 每次都变，但是goroutine闭包会共享相同的
		//变量，导致使用同一个matcher来处理同一个feed
		//为了避免这个问题使用参数传递
		}(matcher, feed)//两个值 传入匿名函数 
		//每次matcher feed 都不一样 所以没用闭包方式访问
		//指针变量方便的在函数内部共享数据，可以让函数访问并修改一个变量的状态
		//这个变量可以在其他函数或goroutine里声明
		//go所有变量都是值传递
		//指针变量指向内存地址，函数间传递指针变量，在传递地址值
	}

	// Launch a goroutine to monitor when all the work is done.
	//一个goroutine来监控是否都完成了
	go func() {//匿名函数 使用闭包访问waitGroup 和results变量
		// Wait for everything to be processed.
		waitGroup.Wait()//递减计数

		// Close the channel to signal to the Display
		//results 之前定义了 用关闭通道的方式 通知display函数 可以退出了
		// function that we can exit the program.
		close(results)
	}()
	//main函数返回那么整个程序就终止了，终止前关闭所有的goroutine

	// Start displaying results as they are available and
	// return after the final result is displayed.
	//调用match的display函数
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
