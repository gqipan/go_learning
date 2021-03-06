package main

import "fmt"

func main() {

	message := make(chan string)
	signal := make(chan bool)

	// 这里是一个非阻塞接收的例子。
	// 如果在 messages 中存在，然后 select 将这个值带入 <-messages case中。
	// 如果不是，就直接到 default 分支中。
	select {
	case msg := <- message:
		fmt.Println("received message: ", msg)
	default:
		fmt.Println("no message received")
	}

	// 一个非阻塞发送的实现方法和上面一样。
	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message: ", msg)
	default:
		fmt.Println("no message sent")
	}

	// 我们可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。
	// 这里我们试图在 messages和 signals 上同时使用非阻塞的接受操作
	select {
	case msg := <- message:
		fmt.Println("received message: ", msg)
	case sig := <-signal:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
