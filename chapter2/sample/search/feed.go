package search
//包 search
import (//引用 库
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"
//定义常量 常量不用指定类型 小写 不对外暴露

// Feed contains information we need to process a feed.
//结构类型 对外暴露Feed
type Feed struct {//包含要处理的袁书记信息
	Name string `json:"site"`//字符串 类型 `` 标记 
	URI  string `json:"link"`
	Type string `json:"type"`
	`` 标记  tag描述json 解码的源数据 用于Feed切片
}

// RetrieveFeeds reads and unmarshals the feed data file.
//search 调用该方法 读取json文件装入到feeds中feed是Feed的切片
func RetrieveFeeds() ([]*Feed, error) {
	//没有参数，会返回两个值  Feed切片 和error 表示是否成功
	// Open the file.
	file, err := os.Open(dataFile)//打开数据文件
	//返回指针，指向File类型的值，第二个返回error类型值
	if err != nil {
		return nil, err
	}
	//返回Feed切片，切片是动态数组的引用类型
	//切片可以操作一组数据，第二个返回值是是不是错误
	// Schedule the file to be closed once
	// the function returns.
	defer file.Close()
	//defer关键字 会安排随后的函数调用在函数返回时执行
	//使用完文件后要主动关闭文件 
	//defer 是在return之后执行关闭cloase以外崩溃也会调用

	// Decode the file into a slice of pointers
	// to Feed values.
	var feeds []*Feed//定义切片 
	//声明 feeds 值是nil 的切片 包含一组指向Feed类型值的指针 *
	//这个切片的每一项是一个指向一个Feed类型的指针
	err = json.NewDecoder(file).Decode(&feeds)
	//json包的NewDecoder函数 返回值上调用Decode方法
	//open返回的文件句柄file调用NewDecoder函数得到一个指向Decoder类型值的的指针
	//之后再调用这个（Decoder）指针的Decode方法，传入切片的地址
	//Decode方法会解码数据 把值以Feed类型值形式存入切片
	//Decode 实现/usr/local/Cellar/go/1.12.9/libexec/src/encoding/json/stream.go:52
	// func (dec *Decoder) Decode(v interface{}) error {
	// Dcode 方法接受一个类型为 interface{}的值作为参考
	// interface 比较特殊 一般会配合 reflect 包里提供反射功能一起使用

	// We don't need to check for errors, the caller can do this.
	return feeds, err
	//调用者会检查错误
}
