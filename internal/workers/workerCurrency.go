package workers

import (
	"context"
	currencyapi "goTest/internal/infrastructure/currencyApi"
	"goTest/internal/modules/currency/storage"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var (
	logger  *zap.Logger
	currStr storage.CurrencyStorage
)

func InitWorkerCurrency(log *zap.Logger, valuteStor storage.CurrencyStorage) {
	c := cron.New(cron.WithSeconds())
	logger = log
	currStr = valuteStor
	log.Info("workerCurrency init")
	_, err := c.AddFunc("0 0 10 * * ?", writeInfoToDb)
	if err != nil {
		log.Fatal("could not schedule task:", zap.Error(err))
	}
	c.Start()
}

func writeInfoToDb() {

	ctx := context.Background()

	// select {
	// case <-ctx.Done():
	// default:
	currencies, err := currencyapi.GetCurrencyByToday()
	if err != nil {
		logger.Error("can not get api", zap.Error(err))
	}

	for _, currency := range currencies {
		err = currStr.InsertCurrency(ctx, currency.NameCurrency, currency.Value)
		if err != nil {
			logger.Error("can not write to db", zap.Error(err))
		}
	}
	logger.Info("make the work")
	// }
}
