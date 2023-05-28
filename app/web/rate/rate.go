package rate

import (
	"encoding/json"
	"log"
	"net/http"
)

type RateInfo struct {
	Time           string
	Asset_id_base  string
	Asset_id_quote string
	Rate           float64
}

func getCoinAPIRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/BTC/UAH", nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.Header.Set("X-CoinAPI-Key", "01F76A3F-1B3A-4B61-84A7-2EF25DD8A3CA")

	return req, nil
}

func parseRateInfo(resp *http.Response) (*RateInfo, error) {
	defer resp.Body.Close()

	var rateInfo RateInfo
	err := json.NewDecoder(resp.Body).Decode(&rateInfo)
	if err != nil {
		return nil, err
	}
	return &rateInfo, nil
}

func GetRate() (float64, error) {
	req, err := getCoinAPIRequest()
	if err != nil {
		return 0, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	rateInfo, err := parseRateInfo(resp)
	if err != nil {
		return 0, err
	}
	return rateInfo.Rate, nil
}