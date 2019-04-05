package type_test

import (
	"testing"
)

type MyInt int64

// 对于类型转换非常严格，必须显示类型转换
func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	// 不支持隐式的类型转换
	//b = a

	b = int64(a)

	var c MyInt
	// 也不支持别名的类型转换
	//c = b

	c = MyInt(b)

	t.Log(a, b, c)

}

func TestPiont(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)

	t.Logf("%T %T", a, aPtr)

	// 不支持指针运算
	//aPtr = aPtr + 1
}

func TestString(t *testing.T) {

	var s string

	t.Log("*" + s + "*")

	// 错误的判断，字符串默认初始化时 “”（空字符串）
	//if s == nil {
	//
	//}
	if s =="" {
		t.Log("s is \"\" ")
	}


}
