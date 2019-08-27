package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

/*
 Go 中最主要的状态管理方式是通过通道间的沟通来完成的，我们在工作池的例子中碰到过，但是还是有一些其他的方法来管理状态的。
 这里我们将看看如何使用 sync/atomic包在多个 Go 协程中进行 原子计数 。
 */
func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for  {
				// 使用 AddUint64 来让计数器自动增加，使用& 语法来给出 ops 的内存地址。
				atomic.AddUint64(&ops, 1)

				// 允许其它 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops: ", opsFinal)
}
