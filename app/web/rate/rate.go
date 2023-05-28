package rate

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
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
		return nil, err
	}

	req.Header.Set("X-CoinAPI-Key", os.Getenv("CoinAPIKey"))

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

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	rateInfo, err := parseRateInfo(resp)
	if err != nil {
		return 0, err
	}
	return rateInfo.Rate, nil
}
