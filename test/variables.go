package main

import (
	"fmt"
	"reflect"
)

func main() {

	//不声明变量类型，自动推断
	var str = "init"
	fmt.Println("str :", str)

	//声明变量类型
	var a, b int = 1, 2
	fmt.Println(b, a)

	var d = true
	fmt.Println(d)

	//声明变量类型不赋值，变量自动初始化为0
	var e int
	fmt.Println(e)

	//声明变量为string时，为空字符串''
	var g string
	//printf函数的%T可以输出变量的类型
	fmt.Printf("%T", g)
	fmt.Println()

	//引入reflect输出g的类型
	fmt.Println(reflect.TypeOf(g))

	//类型自动推断
	f := "short"
	fmt.Println(f)

}
