# gohiveengineSDK
A Software Development Kit for the Hive Blockchain, at least starting with an emphasis on Hive Engine Dex


Building upon the work of cfoxon and contributors to move in the direction it was originally intended. Still currently in the early stages, functionalty isn't robust, but it isn't quite bare-bones either.

 At present it's good for basic use with hive engine, from getting tables(WiP), to creating and cancelling buy/sell orders on Hive Engine.

More functionality to be added as time goes on. For now:

A session creates clients for Hive  and Hive Engine RPCs. If no Url is listed for either/both RPC(s) It will use api.hive.blog / api.hive-engine.com/rpc/ respectively. At present there's no support for private key handling other than a field meant to hold the active wif for use, (while program using this library is running,) is present. 

### Example usage:
To start a new session:
```go
import (
	gohe "github.com/DCowboy/gohiveenginesdk"
)

session := gohe.NewSession(<hive url>, <engine url>, <account name>, <wif string>)
```

Get Hive Engine status:
```go
status, err := session.Status()
//Returns a struct pointer:
chainId := status.ChainId
```

Get Balances for a token:
```go
bxt := session.TokenBalances("BXT")
//Account name is supplied by session
//Returns a struct pointer:
stake := bxt.Stake
```

Get buy/sell books for a token:
```go
book, err :=  session.OrdersBook("buy", "BEE", 10, 0)
// Numbers above are limit and offset, string arguments are case insensitive.
// Returns a struct: book is still an array/slice
buyBook := book.Book
firstPrice := buyBook[0].Price
```

Get account's open orders for a token for a token:
```go
orders, err :=  session.OpenOrders("BEE", 10, 0)
// Numbers above are limit and offset, string arguments are case insensitive.
//Account name is supplied by session
// Returns a struct of each book struct (still returned as a slice)
buyOrders := orders.Buy
firstPrice := buyOrders.Book[0].Price
```
Get trade history for a token:
```go
history, err :=  session.TradesHistory("BEE", 10, 0)
// Numbers above are limit and offset, string arguments are case insensitive.
// Returns a struct of an array of records
log := history.Log
firstRecord := log[0]
firstRecordTimestamp := log[0].Timestamp
```

Get metrics for a token:
```go
book, err :=  session.TokenMetrics("BEE")
//string arguments are case insensitive.
// Returns an array of a struct - Metrics returns as an array because of the query method it uses.
highest := (*response)[0].HighestBid
```

Cancel Hive Engine order:
```go
txid, err := session.CancelOrder("buy", <open order's txid>)
```

Create buy/sell Hive Engine order:
```go
txid, err := session.CreateOrder("sell", "dec", <qty>, <price>)
```

##### Since cfoxon's warning is so good, lets continue it!
-----
WARNING: It is not recommended to stream blocks from public APIs. They are provided as a service to users and saturating them with block requests may (rightfully) result in your IP getting banned

