package main

import (
	"github.com/ethereum/go-ethereum/log"
	"os"
)

func init() {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(log.LvlInfo))
	log.Root().SetHandler(glogger)
}

func main() {

}
