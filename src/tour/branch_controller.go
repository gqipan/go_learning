package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	var sum int = 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("for sum result %v \n", sum)

	// go 中的 while 就是for
	sum2 := 1
	for sum2 < 100 {
		sum2 += sum2
	}
	fmt.Printf("for sum2 result %v \n", sum2)

	// 所以for 可以写无限循环
	//for {
	//	fmt.Printf("Time is %v \n", time.Now())
	//}

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。
	// 除非以 fallthrough 语句结束，否则分支会自动终止。
	// Go 的另一点重要的不同在于 switch 的 case 「无需为常量」，且「取值不必为整数」。
	fmt.Print("Go runs on ")
	switch OS := runtime.GOOS; OS {
	case "darwin":
		fmt.Println("===OS X.")
	case "linux":
		fmt.Println("===Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", OS)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// 无条件的switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("World")
	fmt.Println("Hello")
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if 判断前可新增判断表达式
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}
