package array_test

import "testing"

func TestArrayInit(t *testing.T) {

	// 声明不初始化，默认初始化为0
	var arr [3] int
	t.Log(arr)

	// 下标访问
	arr1 := [4] int{1, 2, 3, 4}
	arr1[1] = 5;
	t.Log(arr1)

	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr2)
}

// 数组的遍历
func TestArrayTravel(t *testing.T) {

	arr3 := [...]int{1, 3, 4, 5}

	// 下标方式遍历
	for i := 0; i < len(arr3); i++ {
		t.Log(i)
	}

	// 类似 foreach， range 返回下标和元素
	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	// 如果不关心下标，可以用 _ 站位符声明
	for _, e := range arr3 {
		t.Log(e)
	}
}

