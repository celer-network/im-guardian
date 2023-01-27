package guardian

import (
	"github.com/celer-network/goutils/log"
)

const (
	FlagMultiChain = "multichain"
	FlagPauser     = "pauser"
)

func Start() {
	done := make(chan bool)
	go StartMonitoring()
	log.Info("app guardian started")
	<-done
}
