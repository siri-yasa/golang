package main

import (
	"fmt"
	"strings"
)

// func WordCount(s string) map[string]int {
// 	return map[string]int{"x": 1}
// }

func WordCount() {
	givenString := "Vamshi is working on golang programming . golang is a programming language"
	split := strings.Split(givenString, " ")

	for i, v := range split {
		fmt.Println(i, v)
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	fmt.Println(m)

}

var m map[string]int

func main() {
	m = make(map[string]int)
	WordCount()
}
