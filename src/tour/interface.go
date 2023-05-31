package main

import (
	"fmt"
)

// 接口

type I interface {
	M()
}

// 声明类型

type T struct {
	S string
}

// 类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字

func (t *T) M() {
	// 即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
	// 注意: 保存了 nil 具体值的接口其自身并不为 nil。
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// 接口值

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	//var i I = T{"hello"}
	//i.M()

	// 接口值
	//接口也是值。它们可以像其它值一样传递。
	//接口值可以用作函数的参数或返回值。
	//在内部，接口值可以看做包含值和具体类型的元组： (value, type)
	//接口值保存了一个具体底层类型的具体值。
	//接口值调用方法时会执行其底层类型的同名方法。

	//var i I
	//i = &T{"Hello"}
	//desc(i)
	//i.M()

	//i = F(math.Pi)
	//desc(i)
	//i.M()

	// 底层值为 nil 的接口值
	//var i I
	//var t *T
	//i = t
	//desc(i)
	//t.M()
	//
	//i = &T{"Hello"}
	//desc(i)
	//i.M()

	// nil 接口值. 接口值也为nil
	//var i I
	//desc(i)
	//i.M() // 如果不调用方法，不会报错

	//指定了零个方法的接口值被称为 *空接口：*
	var i interface{}
	desc2(i)

	i = 41
	desc2(i)

	i = "hello"
	desc2(i)
}

func desc(i I) {
	fmt.Printf("(%v, %T) \n", i, i)
}

func desc2(i interface{}) {
	fmt.Printf("(%v, %T) \n", i, i)
}
