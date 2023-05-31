package main

import "fmt"

// 类型断言 提供了访问接口值底层具体值的方式。
func main() {

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// 若 i 并未保存 T 类型的值，该语句就会触发一个panic。
	// 带上 ok 就不会报错
	x := i.(float64)
	fmt.Println(x)

}
