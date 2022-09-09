package main

import (
	"fmt"
)

var acct map[string] int
var balance map[int]int

acct["Sirisha"] = 101
acct["Vamshi"] = 102
acct["Aswini"] = 103
acct["Kiran"] = 104

balance[101]=300
balance[102]=350
balance[103]=400
balance[104]=450


func deposit(accnum int, tamount int) {
	if _, ok := balance[accnum]; ok {
		balance[accnum] = balance[accnum] + tamount
	} else {
		fmt.Println("Account not found")
	}
}

func withdraw(accnum int, tamount int) {
	if balance[accnum] > tamount {

		balance[accnum] = balance[accnum] - tamount
	} else {
		fmt.Println("Not sufficient funds")
	}
}
func trasfer(withDrawAccount int, depositAmmount int, tamount int) {
	if balance[withDrawAccount] > tamount {

		balance[withDrawAccount] = balance[withDrawAccount] - tamount
		balance[depositAmmount] = balance[depositAmmount] + tamount
	} else {
		fmt.Println("Not sufficient funds")
	}
}
func main() {
	//fmt.Println(a, b, c, d)
	//fmt.Println(w, x, y, z)
	deposit(101, 50)
	fmt.Println()

	withdraw(102, 200)
	fmt.Println()
	trasfer(101, 102, 50)
	fmt.Println()
}
