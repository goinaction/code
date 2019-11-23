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

func threeSlice(){
slice :=[]string{"Apple","Orange","Plum","Bannane","Grape"}
//长度为1个元素 容量为2个元素
newSlice :=slice[2:3:4]
//slice [i:j:k] 第一个表示开始的索引位置，
//第二个表示索引位置加上元素的个数 1 结果是3 为了设置容量，
//从索引位置开始 加上希望容量包含的元素个2  结果是4
//长度j-i 
//容量k-i
//容量大于已有容量会报错 就是超过原来的容量
fmt.Println(newSlice)
//panic: runtime error: slice bounds out of range
//	f :=slice[2:3:6]
//	fmt.Println(f)
}

func sameLens(){
	source := []string{"Apple","Orange","Plum","Bannane","Grape"}
	slice := source[2:3:3]
	//长度和容量一样
	//现在长度和容量一致，还是共享着底层数组但是后面再进行append操作时创建新的数组
	slice=append(slice,"232332")//最佳字符串
	slice=append(slice,"333")//最佳字符串

	s1 :=[]int{1,2}
	s2 :=[]int{3,4}
	fmt.Printf("%v\n",append(s1,s2...))
	for index, value := range slice{
		fmt.Printf("Index:%d value :%s\n",index,value)
	}
	for index ,value := range slice{
		fmt.Printf("value %s vlaue-addr%X elemAddr %X\n",
		value,&value,&slice[index])
		//range 返回的是副本 不是原始值 value 地址不变 要使用index 的方式访问
	}
	for _,val :=range s1{
		fmt.Printf("value%d:",$val)
	}

	for index:=1;index <len(slice);index++{
		fmt.Printf("Index: %d Value :%d\n",index ,slice[index])
	}
}

func twoSlice(){
	slice :=[][]int{{10}, {100,200}}
	slice[0] =append(slice[0],20)
	//append 先增长切片 再将新的整型切片赋值为外发切片的第一个元素
	fmt.Println(slice)
	slice :make([]int, le6)
	newslice := foo(slice)
	fmt.Println(len(newslice),cap(newslice))
}
func foo(slice []int) []int{
	return slice
}//一个切片大小是24个字节  地址8个长度8容量8个字节
