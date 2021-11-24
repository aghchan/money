package main

import (
	controller "github.com/aghchan/money/controller/stock"
	"github.com/aghchan/money/domain/stock"
	"github.com/aghchan/simplegoapp/app"
	"github.com/aghchan/simplegoapp/pkg/twilio"
)

type config struct {
	Twilio struct {
		AccountSid  string `yaml:"account_sid" env:"TWILIO_ACCOUNT_SID"`
		AuthToken   string `yaml:"auth_token" env:"TWILIO_AUTH_TOKEN"`
		PhoneNumber string `yaml:"phone_number" config:"twilio_number"`
	} `yaml:"twilio"`
}

func main() {
	config := &config{}
	routes := []interface{}{
		"/stock/earnings", &controller.StockController{},
		"/test/socket", &controller.StockSocketController{},
	}
	services := []interface{}{
		twilio.NewService,
		stock.NewService,
	}

	server := app.NewApp("localhost", 8081, config, routes, services)

	server.Run()
}
