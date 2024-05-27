package currencyapi

import (
	"encoding/xml"
	"fmt"
	"goTest/internal/models"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

// func GetCurrencyByToday() []models.Currency

type Valute struct {
	XMLName  xml.Name `xml:"Valute"`
	ID       string   `xml:"ID,attr"`
	NumCode  string   `xml:"NumCode"`
	CharCode string   `xml:"CharCode"`
	Nominal  string   `xml:"Nominal"`
	Name     string   `xml:"Name"`
	Value    string   `xml:"Value"`
	Rate     string   `xml:"VunitRate"`
}

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []Valute `xml:"Valute"`
}

func GetCurrencyByNameDate(date, currencyName string) (models.Currency, error) {
	// date = "02/03/2002"

	client := &http.Client{}
	url := "http://cbr.ru/scripts/XML_daily.asp"

	if date != "" {
		url += "?date_req=" + date
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Currency{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36")
	response, err := client.Do(req)
	if err != nil {
		return models.Currency{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.Currency{}, fmt.Errorf("got non-200 status code: %v", err)
	}

	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var v ValCurs
	err = decoder.Decode(&v)
	if err != nil {
		return models.Currency{}, fmt.Errorf("failed to decode XML: %v", err)
	}
	for _, valute := range v.Valute {
		if valute.CharCode == currencyName {
			valute.Value = strings.Replace(valute.Value, ",", ".", -1)
			value, err := strconv.ParseFloat(valute.Value, 64)
			if err != nil {
				return models.Currency{}, fmt.Errorf("failed to parse float: %v", err)
			}
			return models.Currency{NameCurrency: valute.CharCode, Value: value}, nil
		}
	}
	return models.Currency{}, fmt.Errorf("not such Code of Currency")
}

func GetCurrencyByToday() ([]models.Currency, error) {

	client := &http.Client{}
	url := "http://cbr.ru/scripts/XML_daily.asp"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []models.Currency{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36")
	response, err := client.Do(req)
	if err != nil {
		return []models.Currency{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []models.Currency{}, fmt.Errorf("got non-200 status code: %v", err)
	}

	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	var v ValCurs
	err = decoder.Decode(&v)
	if err != nil {
		return []models.Currency{}, fmt.Errorf("failed to decode XML: %v", err)
	}

	currencys := make([]models.Currency, 0)
	for _, valute := range v.Valute {

		valute.Value = strings.Replace(valute.Value, ",", ".", -1)
		value, err := strconv.ParseFloat(valute.Value, 64)
		if err != nil {
			return currencys, fmt.Errorf("failed to parse float: %v", err)
		}

		currencys = append(currencys, models.Currency{
			NameCurrency: valute.CharCode,
			Value:        value,
		})
	}
	return currencys, nil
}
