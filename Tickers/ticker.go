package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select{
			case <-done:
				return
			case ti := <-t.C:
				fmt.Println("Ticked after ",ti)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	done <- true;
	t.Stop()
	fmt.Println("Ticker")
	close(done)

}