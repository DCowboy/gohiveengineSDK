package gohiveenginesdk

import (
	//~ "fmt"
	//~ "log"
	//~ hg "github.com/DCowboy/hivego"
	heg "github.com/DCowboy/hiveenginego"
)

var (
	
)

type Balances struct {
	heg.Balances
}

func (s *Session) AccountName() string {
	return s.account
}

func (s *Session) TokenBalances(symbol string) (*heg.Balances, error) {
	return s.engine.GetBalances(symbol, s.account, 1, 0)
}

func (s *Session) VerifyAcct() (string, error) {
	return s.hive.CheckAccount(s.account)
}

func (s *Session) VerifyKey() (bool, error) {
	return s.hive.CheckKey(s.aKey, s.account)
}
