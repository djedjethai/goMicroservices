package main

import (
	"github.com/djedjethai/bankingLib/logger"
	"github.com/djedjethai/bankingSqlx/app"
	// "log"
)

func main() {

	// log.Println("starting the application")
	logger.Info("Starting the application")
	app.Start()

}
