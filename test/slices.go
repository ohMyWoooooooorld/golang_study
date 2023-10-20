package main

import "fmt"

func main() {
	var s = make([]string, 5)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	var c = make([]string, 5)
	copy(c, s)
	fmt.Println("cpy:", c)
	c = append(c, "d", "e", "f")
	fmt.Println("apd:", c)

	sl1 := c[2:5]
	fmt.Println("sl1", sl1)

	sl2 := c[2:]
	fmt.Println("sl2", sl2)

	sl3 := c[:5]
	fmt.Println("sl3", sl3)

	dlc := []string{"g", "h", "i"}
	fmt.Println("dlc:", dlc)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		//Slice 可以组成多维数据结构。内部的 slice 长度可以不一致，这一点和多维数组不同。
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
