import https from "node:https";

// Tekee pyynnön osoitteeseen, joka annetaan url parametrina.
// Muokkaa myös tekstinä saadun vastauksen JS objektiksi
export default function fetchJson<T>(url: string): Promise<T> {
  const errorMessage = (err: Error) => `Hinnan tarkistus ei onnistunut.\n${err.message}`;

  return new Promise((resolve, reject) => {
    https
      .get(url, (response) => {
        let data = "";

        response.on("data", (chunk) => {
          data += chunk;
        });

        response.on("end", () => {
          try {
            return resolve(JSON.parse(data));
          } catch (err) {
            return reject(errorMessage(err as Error));
          }
        });
      })
      .on("error", (err) => {
        return reject(errorMessage(err));
      });
  });
}
