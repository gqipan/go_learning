package main

import "fmt"

func main() {

	s := []int{2, 3, 5, 7, 11, 13}

	// 截取切片使其长度为 0
	s = s[:0]
	printSlices(s)

	// 拓展其长度
	s = s[:4]
	printSlices(s)

	// 切片的容量是从它的第一个元素开始数，到其「底层数组元素」末尾的个数。
	// 舍弃前两个值. 所以 cap 为4
	s = s[2:]
	printSlices(s)

	s = s[2:]
	printSlices(s)
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
