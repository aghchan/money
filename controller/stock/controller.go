package controller

import (
	"errors"

	"github.com/aghchan/money/domain/stock"
	"github.com/aghchan/simplegoapp/pkg/http"
	"github.com/aghchan/simplegoapp/pkg/logger"
	"github.com/aghchan/simplegoapp/pkg/twilio"
)

type StockController struct {
	StockService stock.Service
	Twilio       twilio.Service
}

func (this StockController) GET(w http.ResponseWriter, req *http.Request) {
	this.StockService.GetStockEarnings()
}

type StockSocketController struct {
	Logger logger.Logger
}

func (this StockSocketController) SOCKET(w http.ResponseWriter, req *http.Request) {
	conn, out, err := http.Upgrade(w, req)
	if err != nil {
		this.Logger.Error(
			"Upgrading to socket",
			"error", err,
		)

		return
	}
	defer conn.Close()
	defer close(out)

	for {
		message, err := http.ReadSocket(conn)
		if err != nil {
			if errors.Is(err, http.ErrUnexpectedSocketClose) {
				this.Logger.Error(
					"Reading from socket",
					"err", err.Error(),
				)
			}

			break
		}

		out <- message
	}
}
