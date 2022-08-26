package main

import (
	"fmt"
)

func main() {
	addResult := addTwoNumbers(10, 20)
	subResult := subTwoNumbers(30, 5)
	mulResult := mulTwoNumber(4, 5)
	divResult := divTwoNumbers(10, 5)
	modResult := modTwoNumbers(11, 3)
	fmt.Println(addResult)
	fmt.Println(subResult)
	fmt.Println(mulResult)
	fmt.Println(divResult)
	fmt.Println(modResult)
}
func addTwoNumbers(x int, y int) int {
	return x + y
}
func subTwoNumbers(a int, b int) int {
	return a - b
}
func mulTwoNumber(s int, j int) int {
	return s * j
}
func divTwoNumbers(c int, d int) int {
	return c / d
}
func modTwoNumbers(g int, v int) int {
	return g % v

}
