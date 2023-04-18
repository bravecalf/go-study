package chain_of_responsibility

import "fmt"

type Doctor struct {
	name string
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.DoctorCheckDone {
		fmt.Println("Patient doctor check already done")
		d.next.execute(p)
		return
	}
	fmt.Printf("doctor [%s] checking patient.\n", d.name)
	p.DoctorCheckDone = true
	if d.next != nil {
		d.next.execute(p)
	} else {
		fmt.Println("Patient Processing Done.")
	}
	return
}

func (d *Doctor) setNext(de Department) {
	d.next = de
}
