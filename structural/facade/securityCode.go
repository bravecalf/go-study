package facade

import "fmt"

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkSecurity(code int) error {
	if s.code != code {
		return fmt.Errorf("security code [%d] is incorrect", code)
	}
	fmt.Println("SecurityCode Verified.")
	return nil
}
