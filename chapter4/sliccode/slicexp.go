package sliccode

func init(){
	fmt.Println("长度和容量都是5个元素:slice :make([]string,5)")
	slice :=make([]string,5)
	slice :=make([]string,3,5)
	fmt.Println("长度3 容量都是5个元素:slice :make([]string,3,5)")
	
	slice := []string{"red","blue","green","yellow","pink"}
	//字符串切片 长度和容量都是5个
	
	slice := []int{10,20,30}//整形切片 长度容量都是3

	slice := []string{99:""}//初始化100个空字符串的元素
	//【】没有指定值创建的是数组
	array :=[3]int{2,3,4}
	
	//nil 切片
	var slice []int
	//空切片
	slice :=make([]int,0)
	slice:=[]int{}
	//调用内部函数 append len cap
}

func setValue(){
	slice := []int{01,32,43,44,44}
	slice[1]=11111
	newSlice := slice[1:3]//只能看4个
	//底层数组【0:01,1:32,2:43,3:44,4:44】 5个
	//切片 后 是 从1个到第3个容量是4个 指针指向第一个 
	//长度容量计算
	//容量是k 的切片 slice [i:j]
	//长度 j-i
	//容量 k-i
	//第一个是表示新切片开始的元素索引位置
	//第二个表示开始的索引位置加上希望包含元素个数 结果索引+个数
	//两个切片共享了底层数组，一个改变另一个也改变
}

func updateValue(){
	slice :=[]int{10,20,30,40,50}
	newSlice := slice[1:3]
	newSlice[1]=999
	newSlice[3]=999//panic: runtime error: index out of range

	fmt.Println(slice)
}

func appendValue(){
	fmt.Println("append",slice)
	slice :=[]int{10,20,30,4,5}
	newSlice:=slice[1:3]
	newSlice =append(newSlice,50)
	fmt.Println("appendlen:",cap(newSlice),"slice",slice,newSlice)
	fmt.Println("增加容量")
	
}