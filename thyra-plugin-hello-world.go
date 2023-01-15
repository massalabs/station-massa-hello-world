package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)

	intSig := make(chan os.Signal, 1)
	signal.Notify(intSig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Fprintf(os.Stderr, "Plugin is initializing.\n")

	for {
		select {
		case <-ticker.C:
			fmt.Fprintf(os.Stdout, "Tic!\n")
		case <-intSig:
			fmt.Fprintf(os.Stderr, "Plugin is shutting down.\nBye!\n")
			return
		}
	}
}
