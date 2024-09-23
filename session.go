package gohiveenginesdk

import (
	//~ "fmt"
	"log"

	hg "github.com/DCowboy/hivego"
	heg "github.com/DCowboy/hiveenginego"
)

var (
	
)

type Session struct {
	hive             *hg.HiveRpcNode
	engine           *heg.HiveEngineRpcNode
	engineActId      string
	account          string
	aKey              string
}

func NewSession(hiveUrl, engineUrl, account, wif string) (*Session, error) {
	urlHive := ""
	urlEngine := ""
	if hiveUrl == "" {
		 urlHive = "https://api.hive.blog"
	} else {
		urlHive = hiveUrl
	}
	if engineUrl == "" {
		 urlEngine = "https://api.hive-engine.com/rpc/"
	} else {
		urlEngine = engineUrl
	}
	engineTest := heg.NewHiveEngineRpc(urlEngine)
	_, eErr := engineTest.GetBalances(symbol, s.account, 1, 0)
	if eErr != nil {
		return nil, eErr
	}
	hiveTest :=  hg.NewHiveRpc(urlHive)
	verify, hErr := hiveTest.checkAccount(&wif)
	if hErr != nil {
		return nil, hErr
	}
	log.Println("verify: %+v", verify) 
	instance := new(Session)
	instance.hive = hiveTest
	instance.engineActId = "mainnet-hive"
	instance.engine = engineTest
	instance.account = account
	instance.aKey = wif
	

	return instance, nil
}

func (s *Session) Status() (*heg.EngineStatus, error) {
	return s.engine.GetStatus()
}

