package main

import (
	"fmt"
	"strconv"
)

func main() {
	inumchk := 21012
	str := strconv.FormatInt(int64(inumchk), 10)
	b := len(str)
	for a := 0; a < len(str); a++ {
		b = b - 1
		if str[a] != str[b] {
			fmt.Println("not palindrome")
			return
		}
	}
	fmt.Println("it's a palindrome")
}
