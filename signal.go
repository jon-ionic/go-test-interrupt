package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a channel to receive the signals.
	sigChan := make(chan os.Signal, 1)

	// Notify the channel on receiving os.Interrupt and syscall.SIGTERM signals.
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Waiting for signals...")

	// Block until a signal is received.
	sig := <-sigChan

	// Print the received signal.
	switch sig {
	case os.Interrupt:
		fmt.Println("Received os.Interrupt signal")
	case syscall.SIGTERM:
		fmt.Println("Received syscall.SIGTERM signal")
	default:
		fmt.Println("Received unknown signal:", sig)
	}

	// Exit the program.
	fmt.Println("Exiting program")
}

