package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main(){
	// Sigchan := make(chan os.Signal, 1)	
	signal.Ignore(syscall.SIGKILL)

	fmt.Printf("the process stated %v", os.Getpid())


	fmt.Printf("signal recieved: %v", sig)
	time.Sleep(4*time.Second)
	fmt.Println("process will be dead")
	select{}


}