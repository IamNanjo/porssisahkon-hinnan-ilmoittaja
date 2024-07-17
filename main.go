package main

import (
	"fmt"
	"time"
)

func main() {
	// Esim. 2024-07-17 13:03:17.584740061 +0300 EEST m=+0.000009098
	now := time.Now().UTC()

	// Esim. 2024-07-17
	today := fmt.Sprintf("%04d-%02d-%02d", now.Year(), now.Month(), now.Day())

	// Esim. 2024-07-17T13:00:00.000Z
	dateString := fmt.Sprintf("%sT%d:00:00.000Z", today, now.Hour())

	latestPrices := fetchLatestPrices()

	if len(latestPrices) < 1 {
		createNotification("Hintatietoja ei löytynyt")
		return
	}

	// Default for float32 is 0
	var currentPrice float32
	var minPrice float32
	var maxPrice float32
	var avgPrice float32

	for _, priceInformation := range latestPrices {
		if minPrice > priceInformation.Price {
			minPrice = priceInformation.Price
		}

		if maxPrice < priceInformation.Price {
			maxPrice = priceInformation.Price
		}

		avgPrice += priceInformation.Price

		if currentPrice == 0 && priceInformation.StartDate == dateString {
			currentPrice = priceInformation.Price
		}
	}

	avgPrice /= float32(len(latestPrices))

	const format string = "2006-01-02T15:04:05.000Z"

	timezone, timeOffset := time.Now().Zone()
	fixedTimezone := time.FixedZone(timezone, timeOffset)

	earliestDate, earliestDateErr := time.Parse(format, latestPrices[len(latestPrices)-1].StartDate)
	latestDate, latestDateErr := time.Parse(format, latestPrices[0].EndDate)

	if earliestDateErr != nil || latestDateErr != nil {
		createNotification("Päivämäärien luku epäonnistui")
	}

	earliestDateLocal := earliestDate.In(fixedTimezone)
	latestDateLocal := latestDate.In(fixedTimezone)

	timespan := fmt.Sprintf("%02d.%02d.%04d %02d:%02d - %02d.%02d.%04d %02d:%02d", 
		earliestDateLocal.Day(),
		earliestDateLocal.Month(), 
		earliestDateLocal.Year(),
		earliestDateLocal.Hour(),
		earliestDateLocal.Minute(),
		latestDateLocal.Day(),
		latestDateLocal.Month(),
		latestDateLocal.Year(),
		latestDateLocal.Hour(),
		latestDateLocal.Minute(),
	)

	createNotification(fmt.Sprintf(`Tämänhetkinen hinta: %.3f €

Hinta aikavälillä %s:
Minimi:     %.3f €
Maksimi:    %.3f €
Keskiarvo:  %.3f €`, currentPrice, timespan, minPrice, maxPrice, avgPrice))

}
