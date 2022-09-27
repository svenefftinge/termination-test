package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println("Received signal")
		fmt.Println(sig)
		secondsIgnored := 0
		for secondsIgnored < 200 {
			time.Sleep(time.Second)
			fmt.Println()
			fmt.Printf("Ignored for %d seconds.", secondsIgnored)
			secondsIgnored++
		}
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
