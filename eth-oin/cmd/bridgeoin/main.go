package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

func init() {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(log.LvlInfo))
	log.Root().SetHandler(glogger)
}

func main() {
	test_eth_client()
}
func test_rpc_client() {
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

func test_rpc_ctx_client() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := rpc.Dial("http://47.242.40.47:6666")
	if err != nil {
		log.Error("dial rpc error", "error", err)
		return
	}

	type Block struct {
		Number string
	}
	var lastBlock Block
	err = client.CallContext(ctx, &lastBlock, "eth_getBlockByNumber", "latest", true)
	if err != nil {
		log.Error("client call error", "error", err)
		return
	}
	log.Info("client call result", "lastBlock", lastBlock)
}

func test_eth_client() {
	client, err := ethclient.Dial("http://47.242.40.47:6666")
	if err != nil {
		log.Error("dial eth error", "error", err)
		return
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Error("get eth error", "error", err)
		return
	}

	fmt.Println(header.Number.String())
}
