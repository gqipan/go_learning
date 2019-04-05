// main 方法的 package 必须是 main
package world

import (
	"fmt"
	"os"
)


func main()  {
	fmt.Println(os.Args)
	if len(os.Args) >1  {
		fmt.Println("hello world!", os.Args[1])
	}
	os.Exit(110)
}