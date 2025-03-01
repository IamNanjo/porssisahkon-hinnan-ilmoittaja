package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// Rajapinnan palauttaman arvon tyyppi
type rawPriceInformation struct {
	Prices []priceInformation `json:"prices"`
}

// Rajapinnan palauttamasta arvosta erotetun prices listan tyyppi
type priceInformation struct {
	Price float32 `json:"price"`
	StartDate string `json:"startDate"`
	EndDate string  `json:"endDate"`
}

// Hakee rajapinnasta viimeisimmät hintatiedot ja palauttaa ne priceInformation listana
func fetchLatestPrices() []priceInformation {
	res, err := http.Get("https://api.porssisahko.net/v1/latest-prices.json")
	if err != nil {
		panic(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		panic(readErr)
	}

	latestPrices := rawPriceInformation{}
	jsonErr := json.Unmarshal(body, &latestPrices)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return latestPrices.Prices
}
