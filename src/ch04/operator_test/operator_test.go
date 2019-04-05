package operator_test

import "testing"

func TestCompareArray(t *testing.T) {

	a := [...] int{1, 2, 3, 4}
	b := [...] int{1, 3, 4, 5}
	//c := [...] int{1, 2, 3, 4, 5}
	d := [...] int{1, 2, 3, 4}
	e := [...] int{1, 3, 2, 4}

	// false 长度相同，元素不同
	t.Log(a == b)

	// 长度不同的数据比较，会编译错误
	//t.Log(a == c)

	// true 长度想同，元素相同，顺序相同
	t.Log(a == d)

	// false 长度相同，元素相同，顺序不同
	t.Log(a == e)

}

const (
	Readable = 1 << iota  //0001
	Writable			  //0010
	Executable			  //0100
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	t.Log(Readable, Writable, Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	// 按位清零, 0111 &^ 0001 = 0110
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	// 0110 &^ 0100 = 0010
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

}
