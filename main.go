package main

import (
	controller "github.com/aghchan/money/controller/stock"
	"github.com/aghchan/money/domain/stock"
	"github.com/aghchan/simplegoapp/app"
	"github.com/aghchan/simplegoapp/pkg/twilio"
)

type config struct {
	Twilio struct {
		AccountSid  string `yml:"account_sid"`
		AuthToken   string `yml:"account_sid"`
		PhoneNumber string `yml:"account_sid" config:"twilio_number"`
	} `yml:"twilio"`
}

func main() {
	config := &config{}
	routes := []interface{}{
		"/stock/earnings", &controller.StockController{},
	}
	services := []interface{}{
		twilio.NewService,
		stock.NewStockService,
	}

	server := app.NewApp("localhost", 8081, config, routes, services)

	server.Run()
}
