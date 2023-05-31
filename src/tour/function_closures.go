package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//pos, neg := adder(), adder()
	f := fib()
	for i := 0; i < 10; i++ {
		//fmt.Println(pos(i), neg(-2*i))
		fmt.Print(f(), ",")
	}
}

func fib() func() int {
	back1, back2 := 0, 1
	return func() int {
		temp := back1
		back1, back2 = back2, back1+back2
		return temp
	}
}
