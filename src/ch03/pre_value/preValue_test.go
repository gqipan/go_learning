package pre_value_test

import (
	"math"
	"testing"
)

func TestPreValue(t *testing.T)  {
	t.Log(math.MaxInt64)
	t.Log(math.MaxInt32)
	t.Log(math.MaxFloat32)
	t.Log(math.MaxFloat64)
	t.Log(math.MaxUint32)
	// 18446744073709551615  overflows int
	//fmt.Println(math.MaxUint64)
	//t.Log(math.MaxUint64)

}


