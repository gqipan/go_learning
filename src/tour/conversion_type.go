package main

import (
	"fmt"
	"math"
)

const pi = 3.14

// go 语言中，类型转换必须是显示转换
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)

	// 类型推导
	var m int = 3
	n := m
	fmt.Printf("n is of type %T \n", n)

	i := 3
	fl := 3.14
	c := 3.14 + 5.12i
	fmt.Printf("var is of type %T  %T  %T  \n", i, fl, c)

	// 常量不能用 := 语法声明。
	const pi2 float64 = 3.14
	const world string = "世界"
	fmt.Printf("var is of type %T  %T %T \n", pi, pi2, world)

	fmt.Printf("var is of type  %T \n", Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)
