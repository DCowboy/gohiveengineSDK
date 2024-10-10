package gohiveenginesdk

import (
	//~ "fmt"
	//~ "log"
	//~ "errors"
	"strings"
	"strconv"
	"encoding/json"

	heg "github.com/DCowboy/hiveenginego"
)

var (
	
)

type OrderBook struct {
	heg.OrderBook
}

type PersonalOrders struct {
	heg.PersonalOrders
}

type History struct {
	heg.History
}

type Metrics struct {
	heg.Metrics
}

func (s *Session) OrdersBook(action, symbol string, limit, offset int) (*heg.OrderBook, error) {
	return s.engine.GetBook(action, symbol, limit, offset)
}

func (s *Session) OpenOrders(symbol string, limit, offset int) (*heg.PersonalOrders, error) {
	return s.engine.GetAccountOrders(symbol, s.account, limit, offset)
}

func (s *Session) TradeHistory(symbol string, limit, offset int) (*heg.History, error) {
	return s.engine.GetHistory(symbol, limit, offset)
}

func (s *Session) TokenMetrics(symbol string) (*heg.Metrics, error) {
	return s.engine.GetMetrics(symbol, 1, 0)
}

func reformatAction(a marketAction) string {
	// marshal struct
	js, _ := json.Marshal(a)
	//change to string
	jstring := string(js)
	// replace exported struct names to fields needed for json
	jstring = strings.Replace(jstring, "ContractName", "contractName", -1)
	jstring = strings.Replace(jstring, "ContractAction", "contractAction", -1)
	jstring = strings.Replace(jstring, "ContractPayload", "contractPayload", -1)
	//replace payload info debending on type
	switch a.ContractPayload.(type) {
		case cancelPayload:
			jstring = strings.Replace(jstring, "Type", "type", -1)
			jstring = strings.Replace(jstring, "Id", "id", -1)
		case tradePayload:
			jstring = strings.Replace(jstring, "Symbol", "symbol", -1)
			jstring = strings.Replace(jstring, "Quantity", "quantity", -1)
			jstring = strings.Replace(jstring, "Price", "price", -1)
	}
	return jstring
}

type marketAction struct {
	ContractName         string
	ContractAction       string
	ContractPayload      interface{}
}
type cancelPayload struct {
	Type                 string
	Id                   string
}

type tradePayload struct {
	Symbol               string
	Quantity             string
	Price                string
}

func (s *Session) CancelOrder (action, txid string) (string, error){
	reqAuths := []string{s.account}
	reqPostAuths := []string{}
	id := string("ssc-" + s.chainId)
	act := marketAction {
		ContractName: "market",
		ContractAction: "cancel",
		ContractPayload: cancelPayload{
			Type: strings.ToLower(action),
			Id: txid,
		},
	}
	
	jstring := reformatAction(act)
	//~ log.Println(jstring)
	tx, err := s.hive.BroadcastJson(reqAuths, reqPostAuths, id, jstring, s.aKey)
	if err != nil {
		return "", err
	}
	//~ log.Println(reqAuths, reqPostAuths, id)
	//~ tx := "debugging"

	return tx, nil
}

func (s *Session) CreateOrder (action, symbol string, qty, price float64) (string, error){
	reqAuths := []string{s.account}
	reqPostAuths := []string{}
	id := string("ssc-" + s.chainId)
	//convert floats to strings
	sQty := strconv.FormatFloat(qty, 'f', 8, 64)
	sPrice := strconv.FormatFloat(price, 'f', 8, 64)
	//set up transaction
	act := marketAction {
		ContractName: "market",
		ContractAction: action,
		ContractPayload: tradePayload{
			Symbol: strings.ToUpper(symbol),
			Quantity: sQty,
			Price: sPrice,
		},
	}
	jstring := reformatAction(act)
	
	//~ fmt.Println(jstring)
	tx, err := s.hive.BroadcastJson(reqAuths, reqPostAuths, id, jstring, s.aKey)
	if err != nil {
		return "", err
	}
	//~ fmt.Println(reqAuths, reqPostAuths, id)
	//~ tx := "debugging"

	return tx, nil
}
