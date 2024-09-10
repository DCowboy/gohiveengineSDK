package gohiveenginesdk

import (
	//~ "fmt"
	//~ "log"
	heg "github.com/DCowboy/hiveenginego"
)

var (
	
)

func (s *Session) AccountName() string {
	return s.account
}

func (s *Session) TokenBalances(symbol string, limit, offset int) (*heg.Balances, error) {
	return s.engine.GetBalances(symbol, s.account, limit, offset)
}
