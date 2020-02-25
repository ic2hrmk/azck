package runner

import (
	"os"
	"os/signal"
	"syscall"
)

func AwaitShutdown() {
	stop := make(chan os.Signal)
	signal.Notify(stop,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	<-stop
}
