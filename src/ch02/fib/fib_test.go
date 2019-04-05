package fib_test

import (
	"fmt"
	"testing"
)


// 或者定义全局变量
var c int
func TestFibList(t *testing.T) {
	//var a int = 1
	//var b int = 1

	// 第二种赋值方式
	//var (
	//	a int = 1
	//	b int = 1
	//)

	// 使用类型推断
	var a = 1
	var b = 1

	// 全局变量赋值
	c = 110

	t.Log(" ", a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log()
}

func TestExchangeVar(t *testing.T)  {

	a := 1
	b := 2

	a, b = b, a

	fmt.Println(a ,b)

}
