package main

import "fmt"

func main() {

	// 数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// 值会初始化为0
	var test_arr [1]int
	fmt.Println(test_arr)

	//每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。
	// 类型 []T 表示一个元素类型为 T 的切片。
	// 切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：
	//a[low : high]
	// 它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

	var s []int = primes[1:4]
	fmt.Println(s)

	// 切片并不存储任何数据，它只是描述了底层数组中的一段。
	// 更改切片的元素会修改其底层数组中对应的元素。
	// 与它共享底层数组的切片都会观测到这些修改。

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	var a_str []string = names[0:2]
	b_str := names[1:3]
	fmt.Println(a_str, b_str)

	b_str[0] = "XXX"
	fmt.Println(a_str, b_str)
	fmt.Println(names)
}
