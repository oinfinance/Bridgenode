package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(log.LvlInfo))
	log.Root().SetHandler(glogger)

	client, err := rpc.Dial("http://47.242.40.47:6666")
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		return
	}

	var account []string
	err = client.Call(&account, "eth_accounts")
	var result string
	//var result hexutil.Big
	err = client.Call(&result, "eth_getBalance", account[0], "latest")
	//err = ec.c.CallContext(ctx, &result, "eth_getBalance", account, "latest")

	if err != nil {
		fmt.Println("client.Call err", err)
		return
	}
	log.Error("test log")

	fmt.Printf("account[0]: %s\nbalance[0]: %s\n", account[0], result)
	log.Info("Got interrupt, shutting down...")
}
