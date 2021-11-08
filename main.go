package main

import (
	controller "github.com/aghchan/money/controller/stock"
	"github.com/aghchan/money/domain/stock"
	"github.com/aghchan/simplegoapp/app"
)

func main() {
	routes := []interface{}{
		"/stock/earnings", &controller.StockController{},
	}
	services := []interface{}{
		stock.NewStockService,
	}

	server := app.NewApp("localhost", 8081, routes, services)

	server.Run()
}
