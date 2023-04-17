package proxy

import "fmt"

// ApplicationProxy 应用代理
// add cache 限制每个用户只能预订一张票
type ApplicationProxy struct {
	application *Application
	cache       map[string]struct{}
}

func (a *ApplicationProxy) bookTicket(cardID string) {
	if _, ok := a.cache[cardID]; ok {
		fmt.Printf("%s has booked ticket.\n", cardID)
		return
	}
	a.application.bookTicket(cardID)
	a.cache[cardID] = struct{}{}
	return
}

func newApplicationProxy() *ApplicationProxy {
	return &ApplicationProxy{
		application: &Application{},
		cache:       make(map[string]struct{}),
	}
}
