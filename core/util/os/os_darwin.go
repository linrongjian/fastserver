package os

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func WaitForSignal() {
	for {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)

		// Block until a signal is received.
		s := <-c
		fmt.Println("Got signal:", s)

		if s == syscall.SIGUSR1 {
			dumpGoRoutinesInfo()
			continue
		}

		if s == syscall.SIGUSR2 {
			reLoadConfig()
			continue
		}

		break
	}
}
