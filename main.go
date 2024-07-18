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

	// Hakee hinnat rajapinnasta
	latestPrices := fetchLatestPrices()

	if len(latestPrices) < 1 {
		createNotification("Hintatietoja ei löytynyt")
		return
	}

	// Valmistele muuttujat hintatiedoille
	// Oletusarvo on 0 tyypille float32
	var currentPrice float32
	var minPrice float32
	var maxPrice float32
	var avgPrice float32

	// Käy hintatiedot läpi ja asettaa pienimmän, suurimman ja nykyisen hinnan.
	// Laskee myös hintojen summan avgPrice muuttujaan, josta lasketaan keskiarvo
	for _, priceInformation := range latestPrices {
		if minPrice > priceInformation.Price {
			minPrice = priceInformation.Price
		}

		if maxPrice < priceInformation.Price {
			maxPrice = priceInformation.Price
		}

		if currentPrice == 0 && priceInformation.StartDate == dateString {
			currentPrice = priceInformation.Price
		}

		avgPrice += priceInformation.Price
	}

	// Viimeistele keskiarvon laskeminen
	avgPrice /= float32(len(latestPrices))

	// ISO-8601 muoto päivämäärälle ja ajalle
	// Vastaa rajapinnan palauttamaa muotoa
	const format string = "2006-01-02T15:04:05.000Z"

	// Paikallinen aikavyöhyke
	timezone, timeOffset := time.Now().Zone()
	fixedTimezone := time.FixedZone(timezone, timeOffset)

	// Lukee aikaisimman ja viimeisimmän päivämäärän ISO-8601 muodosta
	// Tätä käytetään vain timespan muuttujassa, jotta ilmoitukseen saadaan aikaväli
	earliestDate, earliestDateErr := time.Parse(format, latestPrices[len(latestPrices)-1].StartDate)
	latestDate, latestDateErr := time.Parse(format, latestPrices[0].EndDate)

	if earliestDateErr != nil || latestDateErr != nil {
		createNotification("Päivämäärien luku epäonnistui")
	}

	// Muokkaa päivämäärät ja ajat oikeaan aikavyöhykkeeseen
	earliestDateLocal := earliestDate.In(fixedTimezone)
	latestDateLocal := latestDate.In(fixedTimezone)

	// Esim. 18.07.2024 01:00 - 20.07.2024 01:00
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

	// Luo ilmoituksen valmisteltujen muuttujien avulla
	// %s tarkoittaa merkkijonoa (string)
	// %.3f tarkoittaa numeroa 3 desimaalin tarkkuudella
	createNotification(fmt.Sprintf(`Tämänhetkinen hinta: %.3f €
Hinta aikavälillä %s:
Minimi:     %.3f €
Maksimi:    %.3f €
Keskiarvo:  %.3f €`, currentPrice, timespan, minPrice, maxPrice, avgPrice,
	))
}
