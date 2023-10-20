package main

import (
	"fmt"
	"math"
)

//const 声明一个常量
const s string = "constant"

func main() {

	const num = 50000000000
	fmt.Println("string :", s)

	const d = 3e20 / num
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(num))
}
