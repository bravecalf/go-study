package facade

import "fmt"

func TestFacadeDemo() {
	wallet := newWalletFacade("admin", 123456, 88)
	err := wallet.saveMoney("admin", 12345, 10)
	if err != nil {
		fmt.Println(err)
	}
	wallet.queryBalance()

	err = wallet.saveMoney("admin", 123456, 10)
	if err != nil {
		fmt.Println(err)
	}
	wallet.queryBalance()

	err = wallet.spendMoney("admin", 123456, 100)
	if err != nil {
		fmt.Println(err)
	}
	wallet.queryBalance()

	err = wallet.spendMoney("admin", 123456, 80)
	if err != nil {
		fmt.Println(err)
	}
	wallet.queryBalance()
}
