package workers

import (
	"context"
	currencyapi "goTest/internal/infrastructure/currencyApi"
	"goTest/internal/modules/currency/storage"
	"log"
	"time"

	"go.uber.org/zap"
)

var (
	logger  *zap.Logger
	currStr storage.CurrencyStorage
)

func callAt(hour, min, sec int, f func()) error {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	now := time.Now().Local()
	firstCallTime := time.Date(
		now.Year(), now.Month(), now.Day(), hour, min, sec, 0, loc)
	if firstCallTime.Before(now) {
		firstCallTime = firstCallTime.Add(time.Hour * 24)
	}

	duration := firstCallTime.Sub(time.Now().Local())

	log.Println(duration)

	go func() {
		time.Sleep(duration)
		for {
			f()
			time.Sleep(time.Hour * 24)
		}
	}()

	return nil
}

func InitWorkerCurrency(logg *zap.Logger, st storage.CurrencyStorage) {
	logger = logg
	currStr = st
	err := callAt(10, 0, 0, writeInfoToDb)
	if err != nil {
		logger.Fatal("can not to init CurrencyWorker", zap.Error(err))
	}
}

func writeInfoToDb() {

	ctx := context.Background()

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
	logger.Info("made the work")
}
