import notifier from "node-notifier";
import fetchJSON from "./fetch.js";

import type { APILatestPricesResponse } from "./APITypes";

function createNotification(message: string) {
  return notifier.notify({
    title: "Pörssisähkö",
    wait: true,
    message,
  });
}

// Sisältää päivämäärään ja aikaan liittyvät tiedot (tälle päivälle)
const date = new Date();
// Esim. 2024-07-16
const today = `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, "0")}-${date.getDate().toString().padStart(2, "0")}`;
// Esim. 22
const currentHour = `${date.getHours().toString().padStart(2, "0")}`;
// Esim. 2024-07-17T21:00:00:000Z
const dateString = `${today}T${currentHour}:00:00.000Z`;

// Ohjelman suoritus async funktion sisällä
(async () => {
  // Esimerkki funktion palauttamasta arvosta:
  // {
  //   prices: [
  //     {
  //       price: -0.28,
  //       startDate: "2024-07-17T21:00:00:000Z",
  //       endDate: "2024-07-17T22:00:00:000Z",
  //     },
  //     {
  //       price: -0.085,
  //       startDate: "2024-07-17T20:00:00:000Z",
  //       endDate: "2024-07-17T21:00:00:000Z",
  //     },
  //   ],
  // }
  const latestPrices = await fetchJSON<APILatestPricesResponse>(
    "https://api.porssisahko.net/v1/latest-prices.json",
  )
    .then((data) => data.prices)
    .catch((err) => {
      createNotification(err);
      throw "";
    });

  // Suodata vain tälle päivälle olevat hintatiedot
  const currentPrices = latestPrices.filter((price) => price.startDate === dateString);
  // Tämän päivän hintatiedot (jos suodatus onnistui)
  const currentPrice = currentPrices.length > 0 ? currentPrices[0].price : null;

  const earliestDate = new Date(latestPrices[latestPrices.length - 1].startDate);
  const latestDate = new Date(latestPrices[0].endDate);
  // Esim. 17.7.2024 01:00 - 18.7.2024 01:00
  const timeSpan = `${earliestDate.getDate()}.${earliestDate.getMonth() + 1}.${earliestDate.getFullYear()} ${earliestDate.getHours().toString().padStart(2, "0")}:${earliestDate.getMinutes().toString().padStart(2, "0")} - ${latestDate.getDate()}.${latestDate.getMonth() + 1}.${latestDate.getFullYear()} ${latestDate.getHours().toString().padStart(2, "0")}:${latestDate.getMinutes().toString().padStart(2, "0")}`;

  // Lista, joka sisältää vain hinnat
  const pricesOnly = latestPrices.map(({ price }) => price);
  let min = 0;
  let max = 0;
  let avg = 0;

  for (const price of pricesOnly) {
    avg += price;
    if (price < min) min = price;
    if (price > max) max = price;
  }

  avg /= pricesOnly.length;

  // Luo ilmoitus
  createNotification(
    `Tämänhetkinen hinta (${today} ${currentHour}:00):\t ${currentPrice} €

Hinta aikavälillä ${timeSpan}:\t
Minimi:\t\t ${min} €
Maksimi:\t\t ${max} €
Keskiarvo:\t\t ${avg} €`,
  );
})();
