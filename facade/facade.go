package main

import "fmt"

type Account struct {
	accNomor string
	balance  float64
}

func (a *Account) getByAccNomor(accNomor string) Account {
	var account Account
	switch a.accNomor {
	case "001":
		account = Account{"001", 1000}
	case "002":
		account = Account{"002", 2000}
	}

	return account
}

// simplify a complex action/logic/task
type TransactionFacade struct{}

func (t TransactionFacade) transfer(fromAccount string, toAccount string, nominal float64) string {
	acc1 := Account{}
	acc1.getByAccNomor(fromAccount)

	acc2 := Account{}
	acc2.getByAccNomor(toAccount)

	acc1.balance = acc1.balance - nominal
	acc2.balance = acc2.balance + nominal

	return "OK"
}

func main() {
	transactionService := TransactionFacade{}
	status := transactionService.transfer("001", "002", 500)
	fmt.Println(status)
}
