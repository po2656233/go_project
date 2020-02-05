package main

import (
	"github.com/nothollyhigh/kiss/util"
	"go_gate/runner"
	"os"
	"syscall"
)

var Version string
func main() {
	runner.Run(Version)

	//信号退出
	util.HandleSignal(func(sig os.Signal) {
		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			runner.Stop()
			os.Exit(0)
		}
	})
}
