package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"time"
)

type api struct {
	eth  *EthAPI
	near *NearAPI
}

type NearAPI struct {
}

type EthAPI struct {
}

func (e *EthAPI) eth_client() {
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

func (e *EthAPI) eth_rpc_client() {
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

func (e *EthAPI) eth_rpc_ctx_client() {
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
