package main

import (
	"github.com/djedjethai/banking/app"
	"github.com/djedjethai/banking/logger"
	// "log"
)

func main() {

	// log.Println("starting the application")
	logger.Info("Starting the application")
	app.Start()

}
