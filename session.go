package gohiveenginesdk

import (
	//~ "fmt"
	//~ "log"

	hg "github.com/cfoxon/hivego"
	heg "github.com/DCowboy/hiveenginego"
)

var (
	
)

type Session struct {
	hive             *hg.HiveRpcNode
	engine           *heg.HiveEngineRpcNode
	engineActId      string
	account          string
	aKey             string
	Status           *heg.EngineStatus
}

func NewSession(hiveUrl, engineUrl, account, wif string) *Session {
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
	instance := new(Session)
	instance.hive = hg.NewHiveRpc(urlHive)
	instance.engine = heg.NewHiveEngineRpc(urlEngine)
	instance.account = account
	instance.aKey = wif
	instance.Status = instance.engine.GetStatus()
	instance.engineActId = instance.Status.ChainId

	return instance
}

func (s *Session) Status() (*heg.EngineStatus, error) {
	return s.Status
}

