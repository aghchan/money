package controller

import (
	"net/http"

	"github.com/aghchan/money/domain/stock"
	"github.com/aghchan/simplegoapp/pkg/twilio"
)

type StockController struct {
	StockService stock.StockService
	Twilio       twilio.TwilioService
}

func (this StockController) GET(w http.ResponseWriter, req *http.Request) {
	this.StockService.GetStockEarnings()
}
