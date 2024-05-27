package currencyapi

import (
	"testing"
)

func TestGetCurrencyByNameDate(t *testing.T) {
	date := "02/03/2002"
	currencyName := "USD"

	// Вызываете вашу функцию
	result, err := GetCurrencyByNameDate(date, currencyName)

	// Проверяете, что функция не возвращает ошибку
	if err != nil {
		t.Errorf("GetCurrencyByNameDate returned error: %v", err)
	}

	// Проверяете результат вашей функции
	if result.Value <= 0 {
		t.Errorf("GetCurrencyByNameDate returned an invalid value: %v", result)
	}

	if result.Value != 30.9436 {
		t.Errorf("GetCurrencyByNameDate returned an invalid value: %v", result)
	}
}

func TestGetCurrencyByToday(t *testing.T) {

	// Вызываете вашу функцию
	result, err := GetCurrencyByToday()

	// Проверяете, что функция не возвращает ошибку
	if err != nil {
		t.Errorf("GetCurrencyByNameDate returned error: %v", err)
	}

	// Проверяете результат вашей функции
	if result == nil {
		t.Errorf("GetCurrencyByNameDate returned an invalid value: %v", result)
	}
}
