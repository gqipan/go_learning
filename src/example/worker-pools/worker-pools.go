package main

import (
	"fmt"
	"time"
)

/*
  在这个例子中，我们将看到如何使用 Go 协程和通道实现一个工作池 。
 */

func main() {
	// 为了使用 worker 工作池并且收集他们的结果，我们需要2 个通道。
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务。
	for w := 1; w<=3 ; w++ {
		go worker(w, jobs, result)
	}

	// 这里我们发送 9 个 jobs，然后 close 这些通道来表示这些就是所有的任务了。
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a<=9; a++ {
		<-result
	}

	// 执行这个程序，显示 9 个任务被多个 worker 执行。
	// 整个程序处理所有的任务仅执行了 3s 而不是 9s，是因为 3 个 worker是并行的
}

func worker(id int, jobs <- chan int, result chan <- int)  {
	for j := range jobs {
		fmt.Println("worker-", id, "processing job: ", j)
		time.Sleep(time.Second)
		result <- j * 2
	}
}
