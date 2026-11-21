package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main(){
	Sigchan := make(chan os.Signal, 1)	
	signal.Notify(Sigchan, syscall.SIGTERM)

	fmt.Printf("the process stated %v", os.Getpid())

	sig := <-Sigchan

	fmt.Printf("signal recieved: %v", sig)
	time.Sleep(4*time.Second)
	fmt.Println("process will be dead")


}