package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"os"
)

func init() {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(log.LvlInfo))
	log.Root().SetHandler(glogger)
}

func main() {
	//bridge := NewBridge()
	//bridge.Start()
	test_near_client()
	log.Info("service overdddd")
}

// 算数运算请求结构体
type Request struct {
	finality string
}

// 算数运算响应结构体
type Response struct {

}
func test_near_client() {
	// 1) os.StartProcess //
	/*********************/
	/* Linux: */
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("/usr/bin/http", []string{"post", "https://rpc.mainnet.near.org",  "jsonrpc=2.0", "id=dontcare", "method=block", "params:='{\"finality\": \"final\"}'"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)

	//command := "http post https://rpc.mainnet.near.org  jsonrpc=2.0 id=dontcare     method=block     params:='{"finality": "final"}'"
	//cmd := exec.Command("http", "www.baidu.com")
	//output, _ := cmd.CombinedOutput()
	//fmt.Println(string(output))
	//
	//conn, err := jsonrpc.Dial("tcp", "https://rpc.mainnet.near.org")
	//if err != nil {
	//	log.Error("dailing error: ", err)
	//	return
	//}
	//
	//req := Request{"final"}
	//var resp string
	//
	//err = conn.Call("block", req, &resp)
	//if err != nil {
	//	log.Error("arith error: ", err)
	//}
	//log.Info("RRRR",resp)


	//client, err := rpc.Dial("https://rpc.mainnet.near.org")
	//if err != nil {
	//	log.Error("dial rpc error", "error", err)
	//	return
	//}
	//
	//var account []string
	//err = client.Call(&account, "eth_accounts")
	//var result string
	//err = client.Call(&result, "eth_getBalance", account[0], "latest")
	//if err != nil {
	//	log.Error("client call error", "error", err)
	//	return
	//}
	//
	//log.Info("client call result", "account", account[0], "balance", result)
	//

	//command := "http post https://rpc.mainnet.near.org  jsonrpc=2.0 id=dontcare     method=block     params:='{\"finality\": \"final\"}'"
	//cmd := exec.Command("http",command)
	//bytes,err := cmd.Output()
	//if err != nil {
	//	log.Error("err",err)
	//}
	//resp := string(bytes)
	//log.Info("resp",resp)
}
