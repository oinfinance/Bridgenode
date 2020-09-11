package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

func init() {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(log.LvlInfo))
	log.Root().SetHandler(glogger)
}

func main() {

	client, err := rpc.Dial("http://47.242.40.47:6666")
	if err != nil {
		log.Error("dial rpc error", "error", err)
		return
	}

	var account []string
	err = client.Call(&account, "eth_accounts")
	var result string
	err = client.Call(&result, "eth_getBalance", account[0], "latest")
	if err != nil {
		log.Error("client call error", "error", err)
		return
	}

	log.Info("client call result", "account", account[0], "balance", result)
}
