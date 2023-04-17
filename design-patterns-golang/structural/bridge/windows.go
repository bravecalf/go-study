package bridge

import "fmt"

type Windows struct {
	printer Printer
}

func (m *Windows) print() {
	fmt.Println("print request for Windows")
	m.printer.printFile()
}

func (m *Windows) setPrinter(printer Printer) {
	m.printer = printer
}
