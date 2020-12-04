package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Bridge struct {
	wg       sync.WaitGroup
	quit     chan os.Signal
	srv_http *http.Server
}
type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is bridge http server"))
}

func NewBridge() *Bridge {
	bridge := &Bridge{}
	return bridge
}

func (b *Bridge) Start() {
	task := &Task{
		closed: make(chan struct{}),
		ticker: time.NewTicker(time.Second * 10),
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	task.wg.Add(1)
	go func() { defer task.wg.Done(); task.Run() }()

	select {
	case  <-c:
		log.Info("Aborting...\n")
		task.Stop()
	}
	task.wg.Wait()

}

/////////////////////////////////////////////////////

type Task struct {
	closed chan struct{}
	wg     sync.WaitGroup
	ticker *time.Ticker
}

func (t *Task) Run() {
	for {
		select {
		case <-t.closed:
			return
		case <-t.ticker.C:
			getEthBlock()
		}
	}
}

func (t *Task) Stop() {
	close(t.closed)
}

func getEthBlock() {
	client, err := ethclient.Dial("http://47.242.40.47:6666")
	if err != nil {
		log.Error("dial eth error", "error", err)
		return
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Error("get eth HeaderByNumber", "error", err)
		return
	}
	latest := header.Number
	log.Info("result:","latest number is:",latest)

	blockNum := big.NewInt(latest.Int64() -200)
	block, err := client.BlockByNumber(context.Background(),blockNum)
	if err != nil {
		log.Error("get eth BlockByNumber", "error", err)
		return
	}
	trxs := block.Transactions()
	log.Info("block:","num:",block.Number(),"trxs",len(trxs))
	for i, trx := range trxs {
		log.Info("transaction:","Block:",block.Number(),"No:",i," To:",trx.To(),"Amount:",trx.Value())
	}
}

func (b *Bridge) loop() {
	defer b.wg.Done()
}

func (b *Bridge) startHttp() {

	b.quit = make(chan os.Signal)
	signal.Notify(b.quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &Handler{})
	mux.HandleFunc("/test", b.hdlTest)

	b.srv_http = &http.Server{
		Addr:         ":1210",
		WriteTimeout: time.Second * 4,
		Handler:      mux,
	}

	go func() {
		<-b.quit
		if err := b.srv_http.Close(); err != nil {
			log.Warn("Bridge Close Http Server")
		}
	}()

	err := b.srv_http.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Info("Http Server closed")
		} else {
			log.Warn("Http Server closed unexpected:", err)
		}
	}
}

func (b *Bridge) hdlTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown http server"))
	err := b.srv_http.Shutdown(nil)
	if err != nil {
		log.Error("shutdown the server err")
	}
}
