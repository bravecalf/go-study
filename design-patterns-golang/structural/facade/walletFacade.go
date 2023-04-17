package facade

import "fmt"

type WalletFacade struct {
	account      *Account
	securityCode *SecurityCode
	balance      int
}

func newWalletFacade(accountName string, securityCode, money int) *WalletFacade {
	if money < 0 {
		money = 0
	}
	return &WalletFacade{
		account:      newAccount(accountName),
		securityCode: newSecurityCode(securityCode),
		balance:      money,
	}
}

func (w *WalletFacade) saveMoney(accountName string, securityCode, money int) error {
	if err := w.account.checkAccount(accountName); err != nil {
		return err
	}

	if err := w.securityCode.checkSecurity(securityCode); err != nil {
		return err
	}
	w.balance += money
	fmt.Printf("save money [%d] successfully\n", money)
	return nil
}

func (w *WalletFacade) spendMoney(accountName string, securityCode, money int) error {
	if err := w.account.checkAccount(accountName); err != nil {
		return err
	}

	if err := w.securityCode.checkSecurity(securityCode); err != nil {
		return err
	}

	if w.balance < money {
		return fmt.Errorf("balance is not sufficient")
	}

	w.balance -= money
	fmt.Printf("spend money [%d] successfully\n", money)
	return nil
}

func (w *WalletFacade) queryBalance() int {
	fmt.Println("current balance: ", w.balance)
	return w.balance
}
