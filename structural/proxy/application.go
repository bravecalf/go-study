package proxy

import "fmt"

// Application 实际应用
type Application struct {
}

func (a *Application) bookTicket(cardID string) {
	fmt.Printf("%s book ticket successfully.\n", cardID)
}
