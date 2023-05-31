package main

import (
	"fmt"
	"math"
)

// Go 没有类。不过你可以为结构体类型定义方法。
// 方法就是一类带特殊的 接收者 参数的函数。
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
// 在此例中，Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。

type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 记住：方法只是个带接收者参数的函数。
// 现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。

func Abs(v Vertex3) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 为非结构体类型声明方法

type MyFlot float64

func (f MyFlot) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// 指针接受者

func (v *Vertex3) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex3) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 接口

type Abser interface {
	Abs() float64
}

func main() {
	//// 结构体类型方法
	v := Vertex3{3, 4}
	//fmt.Println(v.Abs())
	//fmt.Println(Abs(v))
	//
	//// 非结构体类型方法
	//f := MyFlot(-math.Sqrt2)
	//fmt.Println(f.Abs())

	v.Scale2(10)
	fmt.Println(v.Abs())

	// 指针接受者
	v.Scale(10)
	fmt.Println(v.Abs())

}
