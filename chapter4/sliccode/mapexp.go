package sliccode

//Map 映射是一种数据结构 用于存储一系列无序的键值对
//映射基于键来存储值 基于键快速检索数据
//映射是一个集合 映射是无序的集合
//散列表 包含一组桶
//操作时 都要选择一个桶
// 操作映射时指定的键值传给映射的散列函数 散列函数生成一个索引

func init(){
	//创建一个 键是string 值是int的映射
	dict := make(map[string]int)
	//键值都是string的映射
	dict :=map[string]string{}
	//切片 函数 包含切结构的类型的类型 因为包含引用语义
	//不能作为映射的键
	//dict := map[[]string]int{}
	//声明一个包含字符串切片的映射
	dict :=map[int][]string{}
	
	colors := map[string]string{}
	colors["red"]="#da1337"

	//nil 映射 不能用来存储键值 会报错
	// var colors map[string]string
	// colors["red"] ="da1228"
	
	
}

func isSet(){
	vaule ,exists := colors["blue"]
	if exists {
		fmt.Println(value)
	}
	value := colors["blue"]
	if value !=""{
		fmt.Println(value)
	}
}

func forLoop(){
	colors :=map[string]string{
		"AliceBule" : "#f0f8ff",
		"Coral" : "#ff7F50",
		"DarKray": "#a9a9a9",
		"ForestGreen": "#23b22",
	}
	for key ,value := range colors{
		fmt.Printf("keys %s -value :%s\n",key,value)
	}
	delete(colors,"Coral")
	removeColor(colors,"DarKary")
	for key ,value := range colors{
		fmt.Printf("keys %s -value :%s\n",key,value)
	}
}

func removeColor(colors map[string]string, key string){
	delete(colors,key)
}
