package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	// 使用点号 . 来访问内部属性
	v := Vertex{1, 2}
	x := v.X
	y := v.Y

	fmt.Println(x, y)

	//如果我们有一个指向结构体的指针 p，
	//那么可以通过 (*p).X 来访问其字段 X。
	//不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。

	p := &v
	(*p).X = 1e9
	// 可以直接用 p. 来访问内部
	p.X = 1e9

	fmt.Println(x, v)

	//使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
	var (
		v1 = Vertex{1, 2} //创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1} // Y:0 被隐式地赋予
		v3 = Vertex{}     // X:0 Y:0
		// 特殊的前缀 & 返回一个指向结构体的指针。
		p2 = &Vertex{1, 2}
	)

	fmt.Println(v1, p2, v2, v3)
}
