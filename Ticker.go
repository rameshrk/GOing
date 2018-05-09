package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(1600 * time.Millisecond)

	ticker = time.NewTicker(500 * time.Millisecond)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at second time > ", t)
		}
	}()
	time.Sleep(2001 * time.Millisecond)
	ticker.Stop()

	fmt.Println("Ticker Stopped")
}
