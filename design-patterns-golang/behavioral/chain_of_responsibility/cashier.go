package chain_of_responsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) execute(p *Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Printf("Cashier getting money from patient [%s]\n", p.name)
	return
}

func (c *Cashier) setNext(next Department) {
	c.next = next
}
