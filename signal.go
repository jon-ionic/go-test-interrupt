package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 2)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Waiting for signals...")

	go func() {
		sig := <-sigChan

		switch sig {
		case os.Interrupt:
			fmt.Println("Received os.Interrupt signal")
		case syscall.SIGTERM:
			fmt.Println("Received syscall.SIGTERM signal")
		default:
			fmt.Println("Received unknown signal:", sig)
		}

		fmt.Println("Exiting program")
		os.Exit(0)
	}()

	for {
		fmt.Println("Program running...")
		time.Sleep(10 * time.Second)
	}
}

