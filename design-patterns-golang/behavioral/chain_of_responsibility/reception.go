package chain_of_responsibility

import (
	"fmt"
)

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	if r.next != nil {
		r.next.execute(p)
	} else {
		fmt.Println("Patient Processing Done.")
	}
	return
}

func (r *Reception) setNext(d Department) {
	r.next = d
}

