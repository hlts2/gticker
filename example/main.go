package main

import (
	"fmt"
	"time"

	"github.com/hlts2/gticker"
)

func main() {
	ticker := gticker.NewInstantTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C():
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
