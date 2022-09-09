package main

import (
	"fmt"
)

type Account struct {
	person  string
	accnum  int
	balance int
}

var (
	ali = Account{"siri", 101, 234}
	bunny = Account{"vamshi", 102, 345}
	cherry = Account{"aswini", 103, 567}
	dany = Account{"kiran", 104, 789}
)

func deposit(depositAccount *Account, dAmount int) {
	depositAccount.balance = depositAccount.balance + dAmount
	fmt.Println("Amount ", dAmount, "deposited into account", depositAccount.accnum, " of person ", depositAccount.person, ". Current balance is ", depositAccount.balance)
}

func withdraw(withdrawAccount *Account, wAmount int) {
	withdrawAccount.balance = withdrawAccount.balance - wAmount
	fmt.Println("Amount ", wAmount, "withdraw into account",
		withdrawAccount.accnum, " of person ", withdrawAccount.person, ". Current balance is ", withdrawAccount.balance)
}

func trasfer(withdrawAccount *Account, depositAccount *Account, dAmount int) {
	withdrawAccount.balance = withdrawAccount.balance - dAmount
	depositAccount.balance = depositAccount.balance + dAmount
	fmt.Println("Amount ", dAmount, "withdraw into account", withdrawAccount.accnum, " of person ", withdrawAccount.person, ". Current balance is ", withdrawAccount.balance)
}

func main() {
	fmt.Println(ali, bunny, cherry, dany)

	deposit(&bunny, 50)

	withdraw(&ali, 200)

	trasfer(&cherry, &dany, 100)

}
