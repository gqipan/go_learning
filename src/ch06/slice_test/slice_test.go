package slice_test

import "testing"

// 数组的切片， array[0:len(array)] 左闭右开，避免 outOfIndex
func TestArraySpit(t *testing.T) {

	arr4 := [...] int{1, 2, 3, 4, 5}
	t.Log(arr4)

	// 取前三个元素，如果是从开始，或者到结束可以省略值
	arr4_sec := arr4[:3]
	t.Log(arr4_sec)

	arr4_sec2 := arr4[3:]
	t.Log(arr4_sec2)

	arr4_sec3 := arr4[2:4]
	t.Log(arr4_sec3)
}

func TestSliceInit(t *testing.T) {

}

// 切片的增长
func TestSliceGrowing(t *testing.T) {
	s := [] int{}

	for i := 0; i < 10; i++ {
		// 切片的 cap 会安装 2^n 增长
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

	/** 输出结果，cap 不够就增长
	--- PASS: TestSliceGrowing (0.00s)
	    slice_test.go:33: 1 1
	    slice_test.go:33: 2 2
	    slice_test.go:33: 3 4
	    slice_test.go:33: 4 4
	    slice_test.go:33: 5 8
	    slice_test.go:33: 6 8
	    slice_test.go:33: 7 8
	    slice_test.go:33: 8 8
	    slice_test.go:33: 9 16
	    slice_test.go:33: 10 16
	*/
}

// 切片的共享内存
func TestSliceShareMemory(t *testing.T) {

	// 定义切片, 请注意，没有用 [...]
	year := [] string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	// Jun -> UnKnow
	// 因为共享内存的方式，导致只要修改值，原来数组的值也会被改变
	summer[0] = "UnKnow"
	t.Log(Q2)

	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := [] int{1, 2, 3, 4}
	b := [] int{1, 2, 3, 4}

	// 切片不能直接比较， 只能与 nil 比较
	//if a == b {
	//	t.Log("equal")
	//}

	if a == nil {
		t.Log("equal")
	}

	if b ==nil {
		t.Log("equal")
	}
}
