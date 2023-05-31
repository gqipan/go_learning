package main

import "fmt"

func main() {

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	//q = q[1:len(q)]
	fmt.Println(q)
	// 切片的默认行为
	q = q[1:4]
	fmt.Println(q)
	// 忽略下标，表示从0开始
	q = q[:2]
	fmt.Println(q)
	// 忽略上标，表示到最大值
	q = q[1:]
	fmt.Println(q)

	q = q[:0]
	fmt.Println(q)

	q = q[:4]
	fmt.Println(q)

}
