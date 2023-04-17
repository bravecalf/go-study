package bridge

// Computer 外层抽象
type Computer interface {
	print()
	setPrinter(printer Printer)
}
