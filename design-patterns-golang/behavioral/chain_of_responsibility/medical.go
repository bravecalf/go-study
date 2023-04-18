package chain_of_responsibility

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) execute(p *Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.MedicineDone = true
	if m.next != nil {
		m.next.execute(p)
	} else {
		fmt.Println("Patient Processing Done.")
	}
	return
}

func (m *Medical) setNext(d Department) {
	m.next = d
}
