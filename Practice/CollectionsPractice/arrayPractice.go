package main

import (
	"fmt"
)

func main() {
	//声明数组，必须指定数据类型和存储元素的数量
	var varArray [5]int
	// 声明并初始化
	varArray2 := [5]int{1, 2, 3, 4}
	// 初始化的长度由初始化的值来决定
	varArray3 := [...]int{1, 2, 3, 4, 9, 10}
	fmt.Println(len(varArray3))

	// 通过索引使用数组
	fmt.Println(varArray[2])

	//复制数组，内存空间会复制。编译器会阻止不同类型复制
	varArray4 := varArray2
	varArray4[3] = 20

	fmt.Println(varArray2[3])
	fmt.Println(varArray4[3])

	// 多维数组
	var array = [4][2]int{{10, 11}, {12, 13}}
	fmt.Println(array[1])
	// 可以指定下标
	var array1 = [4][2]int{0: {10, 11}, 1: {12, 13}}
	fmt.Println(array1[1])

	//定义一个长度为5的字符串切片
	slice := make([]string, 5)
	//定义一个长度为3,容量为5的字符串切片，当长度大于容量时，就会自动扩容
	fmt.Println(slice)
	slice1 := make([]string, 3)
	fmt.Println(slice1)
	// 在括号中不写数组则声明的时切片
	slice2 := []string{"123", "345"}
	fmt.Println(slice2)
	// nil切片
	var slice3 []string
	fmt.Println(slice3)
	// 空切片
	slice4 := make([]string, 0)
	fmt.Println(slice4)

	//切片操作，底层共享
	newSlice := []int{1, 2, 3, 4, 5, 6}
	newSlice1 := newSlice[1:3] // 对于容量为k的[i:j]的切片，长度为j-i，容量为k - i
	newSlice1 = append(newSlice1, 60)
	fmt.Println(newSlice)  // [1 2 3 60 5 6]
	fmt.Println(newSlice1) //[2 3 60]

	newSlice2 := newSlice[1:3:3]
	newSlice2 = append(newSlice2, 70)
	fmt.Println(newSlice)  // [1 2 3 60 5 6]
	fmt.Println(newSlice2) // [2 3 70]

	newSlice3 := append(newSlice1, newSlice2...)
	fmt.Println(newSlice3) // [2 3 60 2 3 70]

	for index, value := range newSlice3 {
		fmt.Printf("index is %d, value is %d, ElemAdrr is %X, ValueAdrr is %X \n", index, value, &newSlice3[index], &value)
	}

	/**
	index is 0, value is 2, ElemAdrr is C000010230, ValueAdrr is C0000141A0
	index is 1, value is 3, ElemAdrr is C000010238, ValueAdrr is C0000141A0
	index is 2, value is 60, ElemAdrr is C000010240, ValueAdrr is C0000141A0
	index is 3, value is 2, ElemAdrr is C000010248, ValueAdrr is C0000141A0
	index is 4, value is 3, ElemAdrr is C000010250, ValueAdrr is C0000141A0
	index is 5, value is 70, ElemAdrr is C000010258, ValueAdrr is C0000141A0
	*/

	var capSlice []int
	for i := 0; i < 511; i++ {
		capSlice = append(capSlice, i)
	}
	fmt.Println(cap(capSlice)) // 512
	for i := 0; i < 2; i++ {
		capSlice = append(capSlice, i)
	}
	fmt.Println(cap(capSlice)) // 1024
	for i := 0; i < 525; i++ {
		capSlice = append(capSlice, i)
	}
	fmt.Println(cap(capSlice)) // 1280

	// 创建映射，键的类型是string，值的类型是int
	dict := make(map[string]int)
	dict1 := map[string]int{"red": 1, "blue": 2}

	value, exists := dict1["red"]   // 第一个返回值表示值，第二个返回值表示这个键是否存在
	value1, exists1 := dict1["eee"] // 如果不存在键，value返回默认零值
	value2 := dict1["red"]          // 只有一个参数的时候返回值
	fmt.Println(exists, value)      //true 1
	fmt.Println(exists1, value1)    //false 0
	fmt.Println(value2)             // 1
	fmt.Println(dict)

	//使用内置delete()删除值
	delete(dict1, "red")

	//迭代映射
	for key, value := range dict1 {
		fmt.Printf("key is %s, value is %d\n", key, value)
	}
	/**
	key is blue, value is 2
	*/

}
