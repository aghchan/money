package controller

import (
	"net/http"

	"github.com/aghchan/money/domain/stock"
	"github.com/aghchan/simplegoapp/pkg/twilio"
)

type StockController struct {
	StockService stock.Service
	Twilio       twilio.Service
}

func (this StockController) GET(w http.ResponseWriter, req *http.Request) {
	this.StockService.GetStockEarnings()
}
