package main

import "fmt"

func main() {
	//初始化一个数组 [个数]类型
	var arr [5]int
	fmt.Println(arr)

	arr1 := [5]int{1, 2, 3, 4, 5}
	arr1[3] = 0
	fmt.Println(arr1)

	fmt.Println(len(arr1))

	var twoArray [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoArray[i][j] = i + j
		}
	}
	fmt.Println(twoArray)
}
