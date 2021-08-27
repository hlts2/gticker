# gticker

gticker is time ticker with immediate first tick.


The Go standard library has a really cool type – Ticker.
Tickers are used when you want to do something at a regular interval.
However, if we want to tic immediately the first time, we can cut it out to a function and call it before tic.
This package was created for the purpose of being able to handle a different approach from that implementation.

## Requirement

Go 1.16

## Installation

```shell
go get github.com/hlts2/gticker
```

## Example

```go

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
```

## Author
[hlts2](https://github.com/hlts2)

