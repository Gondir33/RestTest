package currencyapi

import (
	"encoding/xml"
	"fmt"
	"goTest/internal/models"
	"net/http"
)

// func GetCurrencyByNameDate(date, currencyName string) string {
// 	// date = "02/03/2002"
// 	response, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp?date_req=" + date)
// 	if err != nil {
// 		fmt.Printf("The HTTP request failed with error: %sn", err)
// 	}

// 	data, _ := io.ReadAll(response.Body)
// 	return string(data)
// }

func GetCurrencyByToday() []models.Currency

type ValCurs struct {
	Valute []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  string `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func GetCurrencyByNameDate(date, currencyName string) string {
	// 	// date = "02/03/2002"
	response, err := http.Get("http://www.cbr.ru/scripts/XML_daily.asp?date_req=" + date)
	if err != nil {
		fmt.Printf("The HTTP request failed with error: %sn", err)
	}
	defer response.Body.Close()

	var valCurs ValCurs
	err = xml.NewDecoder(response.Body).Decode(&valCurs)
	if err != nil {
		fmt.Printf("Failed to decode XML: %sn", err)
	}

	for _, valute := range valCurs.Valute {
		if valute.Name == currencyName {
			return valute.Value
		}
	}
	return "Currency not found"
}
