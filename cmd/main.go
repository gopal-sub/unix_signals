package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)


func main(){
    signal.Ignore(syscall.SIGTERM)

    fmt.Println("PID:", os.Getpid())
    fmt.Println("SIGTERM is being ignored")

    select {}


}