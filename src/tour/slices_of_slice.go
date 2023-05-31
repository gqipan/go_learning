package main

import (
	"fmt"
	"strings"
)

// 切片可包含任何类型，甚至包括其它的切片。
func main() {

	//// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[0][2] = "X"
	board[1][0] = "O"
	board[1][2] = "X"
	board[2][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i], " "))
	}

}
