package gohiveenginesdk

import (
	"fmt"
	"log"

	hg "github.com/DCowboy/hivego"
	heg "github.com/DCowboy/hiveenginego"
)

var (
	
)

type Session struct {
	hive             *hg.HiveRpcNode
	engine           *heg.HiveEngineRpcNode
	chainId          string
	account          string
	aKey             *string
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
	status, eErr := engineTest.GetStatus()
	log.Printf("Hive Engine Status: %+v", *status)
	if eErr != nil {
		return nil, fmt.Errorf("Engine Test err: %s", eErr)
	}
	hiveTest :=  hg.NewHiveRpc(urlHive)
	verified, hErr := hiveTest.CheckAccount(account)
	if hErr != nil {
		return nil, fmt.Errorf("Hive Test err: %s", hErr)
	}
	_, kErr := hiveTest.CheckKey(&wif, account)
	if kErr != nil {
		return nil, kErr
	}
	instance := new(Session)
	instance.hive = hiveTest
	instance.chainId = status.ChainId
	instance.engine = engineTest
	instance.account = verified
	instance.aKey = &wif
	

	return instance, nil
}

func (s *Session) Status() (*heg.EngineStatus, error) {
	return s.engine.GetStatus()
}

