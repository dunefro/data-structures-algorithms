package main

import "fmt"

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	sc := &securityCode{
		code: code,
	}
	return sc
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security code is incorrect")
	}
	fmt.Println("Account verified")
	return nil
}
