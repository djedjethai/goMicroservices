package main

import (
	"github.com/djedjethai/bankingSqlx/app"
	"github.com/djedjethai/bankingSqlx/logger"
	// "log"
)

func main() {

	// log.Println("starting the application")
	logger.Info("Starting the application")
	app.Start()

}
