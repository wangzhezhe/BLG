package BasicFunctions

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func installSignalHandler() {
	c := make(chan os.Signal, 1)
	//add more types of signal here
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	go func() {
		sig := <-c
		log.Println("get the signal and exit")
		log.Printf("Exiting given signal: %v", sig)
		//do some operations
		os.Exit(0)
	}()
}
