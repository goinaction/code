package arraycode


func init(){
	var array0 [5]int
	fmt.Println(" var array0 [5]int 声明一个5个元素的int",array0)
	array1 := [5]int{10,20,30,40,50}
	fmt.Println(" array1 := [5]int{10,20,30,40,50} 声明一个5个元素的int",array1)
	array3 := [...]int{11,21,31,41,51}
	fmt.Println(" array3 := [...]int{11,21,31,41,51} 使用...代替数量 声明一个5个元素的int",array3)
	array4 := [5]int{1:10, 2:20}
	fmt.Println(" array4 := [5]int{1:10, 2:20} //声明一个有5个元素的数组 初始化1，2为索引的元素 其余是0 ")
	for index,val :=range array4{
		fmt.Println("array4:index->",index,"-",val)
	}
	array0[1]=23
	fmt.Println("array0 使用【】访问索引操作array0[1]=23 ",array0)
	var pointstring = "声明一个所有元素都是指针的数组。使用*运算符访问元素指针所指向的值\n"
	array5 := [5]*int{0:new(int),1:new(int)}
	var code5="array5 := [5]*int{0:new(int),1:new(int)}//声明包含5个元素的指整数的数组，
	/ 用整数型指针初始化索引为0 和1 的数组元素 *int 类型"
	*array[0]=10
	*array[1]=20
	var code5_1=" *array[0]=10 *array[1]=20 初始化索引0 1  \n"
	fmt.Println(pointstring,array5,code5,code5_1)
	array1=array3
	fmt.Println("array1=array3 把array3赋值array1 两个必须长度类型一样 array1:",array1)

	
	fmt.Println("把一个指针数组赋值给另一个")
	var array6 [3]*string
	array7 := [3]*string{new(string),new(string),new(string)}
	*array7[0]="red"
	*array7[1]="blue"
	*array7[2]="green"
	array6=array7
	//两个数组指向同一个字符串
	//如果改变 *array7[2] 那么*array6[2]的值也会改变
	var array8[4][2]int
	array:=[4][2]int{{0,0},{1,1},{2,2},{3,3},{4,4}}
	array:=[4][2]int{{0,0},{3,3},3:{4,4}}
	//位置必须是一个对应一个 尤其是加了索引不可以越位
	//从前向后数数如果空就是0 
	array:=[2][2]int{1:{0:3}}
	//访问二维数组
	var array[2][2]int
	array[0][0] =10
	array[0][1]=23
	array[1][0]=44

	var array9 [2]int =array1[1]
	var value int = array[0][1]

}	
var array [le6]int //8M的数组
foo(array)
func foo(array [le6]int){
	fmt.Println(array)
}//每次传递都要在栈上分配8M的内存整个数组的值8m的内存在复制

var arrayle5[le6]int
foo6(&array)
func foo6(array *[le6]int){
	//传入地址 在栈上分配了8个字节的内存
	//
}