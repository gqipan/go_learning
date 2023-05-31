package main

import "fmt"

var rangePow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range rangePow {
		fmt.Printf("2**%d = %d \n", i, v)
	}

	// 用下划线，忽略取值
	for i, _ := range rangePow {
		fmt.Printf("2**%d \n", i)
	}

	for _, value := range rangePow {
		fmt.Printf("2** = %d \n", value)
	}
	// 如果是第二个值，甚至可以不用下划线来忽略
	for i := range rangePow {
		fmt.Printf("2**%d \n", i)
	}

	pow3 := make([]int, 10)
	fmt.Printf("pow3 ******** %v \n", pow3)
	for i := range pow3 {
		pow3[i] = 1 << uint(i)
	}
	for _, value := range pow3 {
		fmt.Printf("%d \n", value)
	}
}
