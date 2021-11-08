package controller

import (
	"net/http"

	"github.com/aghchan/money/domain/stock"
)

type StockController struct {
	StockService stock.StockService
}

func (this StockController) GET(w http.ResponseWriter, req *http.Request) {
	this.StockService.GetStockEarnings()
}
