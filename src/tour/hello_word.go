package main

// 可单独导入包
import (
	"math"
	"time"
)

// 也可分组导入, 用括号代替
import (
	"fmt"
	"math/rand"
	"os"
)

// var c, python, java bool
// 如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
var c, python, java = true, false, "no!"

func main() {
	fmt.Println("Helle World!")
	fmt.Println("The time is ", time.Now())
	fmt.Println("My favorite number is ", rand.Int())
	fmt.Printf("Now you have  %g problems\n", math.Sqrt(7))
	//只能引用已导出的包， 大写开头的表示为已经导出的包
	fmt.Println(math.Pi)
	fmt.Println("Test invoke add func result", add(3, 2))
	fmt.Println("Test invoke add2 func result", add(42, 13))
	// 测试多值返回
	a, b := swap("Hello", "World")
	fmt.Println("Test return mutil value", a, b)
	//直接返回
	fmt.Println(split(17))

	//var 语句可以出现在包或函数级别。
	//var i int //默认值是0
	// 带初始值，可省略
	var i, j int = 1, 2
	fmt.Println(i, j, c, python, java)

	//在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
	//函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
	m := 1
	fmt.Println("在函数中，使用 := 代替 var 定义。", m)

	if len(os.Args) > 1 {
		fmt.Println("Try get os args", os.Args)
	}
}

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

// 多值返回
func swap(x, y string) (string, string) {
	return y, x
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
// 返回值的名称应当具有一定的意义，它可以作为文档使用。
// 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
