package facade

import "fmt"

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account Name [%s] is incorrect", accountName)
	}
	fmt.Println("Account Verified.")
	return nil
}
