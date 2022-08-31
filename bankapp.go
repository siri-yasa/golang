package main

import "fmt"

var acct map[string]int
var balance map[int]int

func main() {

	acct = make(map[string]int)
	balance = make(map[int]int)

	acct["Sirisha"] = 101
	acct["Vamshi"] = 102
	acct["Aswini"] = 103
	acct["Kiran"] = 104

	balance[101] = 1000
	balance[102] = 2000
	balance[103] = 3000
	balance[104] = 4000

	deposit(101, 50)
	fmt.Println(balance[101])

	trasfer(101, 104, 50)
	fmt.Println(balance[101], balance[104])
}

func deposit(accnum int, amount int) {
	if _, ok := balance[accnum]; ok {
		balance[accnum] = balance[accnum] + amount
	} else {
		fmt.Println("Account not found")
	}
}

func withdraw(accnum int, amount int) {
	if balance[accnum] > amount {

		balance[accnum] = balance[accnum] - amount
	} else {
		fmt.Println("Not sufficient funds")
	}
}
func trasfer(withDrawAccount int, depositAmmount int, amount int) {
	if balance[withDrawAccount] > amount {

		balance[withDrawAccount] = balance[withDrawAccount] - amount
		balance[depositAmmount] = balance[depositAmmount] + amount
	} else {
		fmt.Println("Not sufficient funds")
	}
}
