package stock

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/Finnhub-Stock-API/finnhub-go/v2"
	"go.uber.org/zap"
)

type StockService interface {
	GetStockEarnings() error
}

func NewStockService(logger *zap.SugaredLogger) StockService {
	ss := stockService{
		logger:            logger,
		stockEarningsSeen: make(map[string]bool),
	}

	go func(ss *stockService) {
		ss.logger.Infow("started up the earnings poller")

		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			ss.GetStockEarnings()
		}
	}(&ss)

	return &ss
}

type stockService struct {
	logger *zap.SugaredLogger

	stockEarningsSeen map[string]bool
}

func (this stockService) GetStockEarnings() error {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", "c64cv62ad3i8bn4fj430")
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi
	now := time.Now().UTC()

	earningsCalendar, _, err := finnhubClient.EarningsCalendar(context.Background()).From(now.String()).To(now.String()).Execute()
	if err != nil || earningsCalendar.EarningsCalendar == nil {
		if err == nil {
			err = errors.New("missing earnings calendar data")
		}

		this.logger.Errorw(
			"Getting earnings calendar",
			"err", err,
		)

		return err
	}

	found := false

	for _, stock := range *earningsCalendar.EarningsCalendar {
		earningsDate, err := time.Parse("2006-01-02", *stock.Date)
		if err != nil {
			this.logger.Errorw(
				"Parsing earnings date",
				"err", err,
				"stock", stock,
			)
		}

		if stock.Symbol == nil ||
			earningsDate.Day() != time.Now().UTC().Day() {
			continue
		}

		if _, seen := this.stockEarningsSeen[*stock.Symbol]; !seen &&
			stock.EpsActual != nil && stock.EpsEstimate != nil &&
			stock.RevenueActual != nil && stock.RevenueEstimate != nil &&
			*stock.RevenueActual > 0.0 && *stock.RevenueEstimate > 0.0 &&
			stock.GetEpsActual() >= (stock.GetEpsEstimate()*1.1) &&
			stock.GetRevenueActual() >= (stock.GetRevenueEstimate()*1.1) {

			this.logger.Warnw(
				"Checkout this stock",
				"symbol", stock.Symbol,
				"stock", stock,
			)

			this.stockEarningsSeen[*stock.Symbol] = true
			found = true
		}
		// TODO: 37 minute delay? 1636407366.298318 need about 5 min or so on LOTZ
	}

	if !found {
		this.logger.Debugw(
			"No stocks worth buying found",
		)
	}

	return nil
}

func getQuarter() int {
	month, _ := strconv.Atoi(time.Now().Month().String())
	return ((month + 2) / 3)
}
