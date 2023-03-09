package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/raven0520/wallet/app"
	"github.com/raven0520/wallet/router"
)

var env = "./config/"

func main() {
	err := app.InitModule(env, []string{"base"})
	if err != nil {
		fmt.Printf("Failed Init %s", err) // Output err message to console
		return
	}
	if err = app.InitConsulServer(); err != nil {
		fmt.Println(err)
		return
	}
	router.HTTPServerStart() // Start Http Service
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	router.HTTPServerStop() // Stop Http Service
	close(exit)
}
