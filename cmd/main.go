package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main(){
    signal.Ignore(syscall.SIGTERM)

    fmt.Println("PID:", os.Getpid())
    fmt.Println("SIGTERM is being ignored")

	time.Sleep(5*time.Minute)


}