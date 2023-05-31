package main

import (
	"fmt"
	"strings"
)

func main() {

	m := make(map[string]int)

	m["AW"] = 42
	fmt.Println("The value:", m["AW"])

	m["AW"] = 48
	fmt.Println("The value:", m["AW"])

	delete(m, "AW")
	fmt.Println("The value:", m["AW"])

	v, ok := m["AW"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	fields := strings.Fields(s)
	for _, value := range fields {
		_, ok := m[value]
		if ok {
			m[value] = m[value] + 1
		} else {
			m[value] = 1
		}
	}
	return m
}
