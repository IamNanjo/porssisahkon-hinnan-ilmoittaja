# Pörssisähkön hinnanseuranta

`This program is meant to be used to track and notify the user of electricity prices. It uses the API for a Finnish service, therefore the documentation and notifications will be in Finnish for this project.`

Tämä ohjelma näyttää ilmoituksen, jossa näkyy tämänhetkinen pörssisähkön hinta.
Voit ladata suoritettavan tiedoston [GitHub Releases sivulta](https://github.com/IamNanjo/porssisahkon-hinnan-ilmoittaja/releases/latest)

Voit suorittaa koodin myös omalla [Bun](https://bun.sh) tai [Node.js](https://nodejs.org) asennuksellasi seuraavien komentojen avulla

```bash
# Node.js
npm install
npm run build
node index.js

# Bun
bun install
bun index.ts
```

Voit luoda suoritettavan tiedoston [Bun](https://bun.sh) avulla

Voit hyödyntää tähän [build.sh](./build.sh) (linux) tiedostoa, tai voit suorittaa seuraavan komennon itse koodin hakemistossa, jonka jälkeen löydät suoritettavan tiedoston `dist` kansiosta.
Suoritettavan tiedoston luominen Windows järjestelmissä vaatii enemmän työtä, koska Bun ei toimi kovinkaan hyvin Windows järjestelmissä.
Muut työkalut, kuten [pkg](https://www.npmjs.com/package/pkg) toimivat, mutta ne vaativat manuaalista tiedostojen siirtoa, koska kaikkia tarvittavia tiedostoja ei voi sisällyttää .exe tiedostoon.

```bash
bun build --compile index.ts --outfile dist/hinta-tarkkailija
```
