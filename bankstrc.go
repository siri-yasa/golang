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
	a = Account{"siri", 101, 234}
	b = Account{"vamshi", 102, 345}
	c = Account{"aswini", 103, 567}
	d = Account{"kiran", 104, 789}
	//p = &Account{"lord", 100, 480} //has type of *Account
)

/*func (p *Account) deposit(dAmount int) {
	p.balance = p.balance + dAmount
}*/

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
	fmt.Println(a, b, c, d)

	deposit(&b, 50)

	withdraw(&a, 200)

	trasfer(&c, &d, 100)

}
