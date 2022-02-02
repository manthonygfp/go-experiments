package main

import (
	"fmt"
	"time"
)

func sleepAndWake(c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}

func main() {
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)

	go sleepAndWake(done)

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			ticker.Reset(time.Second)
			go sleepAndWake(done)
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}
