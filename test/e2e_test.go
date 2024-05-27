package test

import (
	"fmt"
	"goTest/config"
	"goTest/internal/infrastructure/component"
	"goTest/internal/infrastructure/godecoder"
	"goTest/internal/infrastructure/responder"
	"goTest/internal/modules"
	mock_storage "goTest/internal/modules/currency/storage/mocks"
	"goTest/internal/storages"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type inputData struct {
	date string
	val  string
}

func TestCurrency_GetCurrency(t *testing.T) {
	logger, _ := zap.NewProduction()
	decoder := godecoder.NewDecoder()
	responseManager := responder.NewResponder(decoder, logger)
	components := component.NewComponents(config.AppConf{}, responseManager, decoder, logger)

	// inti mock storages
	mockStorage := mock_storage.NewMockCurrencyStorage(gomock.NewController(t))
	newStorages := storages.NewStorages(mockStorage)

	services := modules.NewServices(newStorages, components)

	controllers := modules.NewControllers(services, components)

	//init server
	r := chi.NewRouter()
	r.Get("/currency", controllers.Currency.GetCurrency)

	tests := []struct {
		name   string
		input  inputData
		output string
	}{{
		name: "good request",
		input: inputData{
			date: "02/03/2002",
			val:  "USD",
		},
		output: "{\"model\":\"USD\",\"value\":30.9436}\n",
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			target := fmt.Sprintf("/currency?date=%s&val=%s", test.input.date, test.input.val)
			req := httptest.NewRequest("GET", target, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, test.output, w.Body.String())

		})
	}
}
