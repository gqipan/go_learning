package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func main() {
	m = make(map[string]Vertex2)
	m["BL"] = Vertex2{
		40.65,
		-74.399,
	}
	fmt.Println(m["BL"])
	fmt.Println(m2)
}

var m2 = map[string]Vertex2{
	"BL": Vertex2{
		1, 2,
	},
	"GG": Vertex2{
		2, 4,
	},
}
