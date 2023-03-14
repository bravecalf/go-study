package bridge

func TestBridgeDemo() {
	hp := HP{}
	canon := Canon{}
	mac := Mac{
		printer: hp,
	}
	mac.print()

	mac.setPrinter(canon)
	mac.print()

	windows := Windows{}
	windows.setPrinter(hp)
	windows.print()
}
