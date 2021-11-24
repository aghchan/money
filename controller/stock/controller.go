package controller

import (
	"github.com/aghchan/money/domain/stock"
	"github.com/aghchan/simplegoapp/pkg/http"
	"github.com/aghchan/simplegoapp/pkg/twilio"
)

type StockController struct {
	http.Controller

	StockService stock.Service
	Twilio       twilio.Service
}

func (this StockController) GET(w http.ResponseWriter, req *http.Request) {
	this.StockService.GetStockEarnings()
}

type StockSocketController struct {
	http.Controller
}

func (this StockSocketController) SOCKET(w http.ResponseWriter, req *http.Request) {
	conn, out, err := this.Upgrade(w, req)
	if err != nil {
		return
	}
	defer conn.Close()
	defer close(out)

	for {
		_, err := this.ReadSocket(conn)
		if err != nil {
			break
		}

		err = this.SendMessage(out, response{Sample: "testing"})
		if err != nil {
			break
		}
	}
}

type response struct {
	Sample string `json:"myTest"`
}
