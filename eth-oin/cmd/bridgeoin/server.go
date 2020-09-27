package main

import (
	"github.com/ethereum/go-ethereum/log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type bridge struct {
	wg       sync.WaitGroup
	quit     chan os.Signal
	srv_http *http.Server
}
type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is bridge http server"))
}

func (b *bridge) Start() {

}

func (b *bridge) loop() {
	defer b.wg.Done()
}

func (b *bridge) startHttp() {

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

func (b *bridge) hdlTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,shutdown http server"))
	err := b.srv_http.Shutdown(nil)
	if err != nil {
		log.Error("shutdown the server err")
	}
}
