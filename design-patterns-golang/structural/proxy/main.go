package proxy

func TestProxyDemo() {
	ap := newApplicationProxy()
	ap.bookTicket("12345678")
	ap.bookTicket("12345678")
	ap.bookTicket("1234")
}
