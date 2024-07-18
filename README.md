# Pörssisähkön hinnanseuranta

`This program is meant to be used to track and notify the user of electricity prices. It uses the API for a Finnish service, therefore the documentation and notifications will be in Finnish for this project.`

Tämä ohjelma näyttää ilmoituksen, jossa näkyy tämänhetkinen pörssisähkön hinta.
Voit ladata suoritettavan tiedoston [GitHub Releases sivulta](https://github.com/IamNanjo/porssisahkon-hinnan-ilmoittaja/releases/latest)

Voit muokata ohjelmaa ja luoda itse oman suoritettavan tiedoston [Go](https://go.dev/):n avulla.

Lataa vaaditut moduulit ja luo suoritettava tiedosto seuraavilla komennoilla

```bash
# Lataa vaaditut moduulit
go get

# Suorita ohjelma testataksesi toiminnan
go run .

# Luo suoritettava tiedosto
go build -o dist/
```

Löydät suoritettavan tiedoston dist/ kansion sisältä tämän jälkeen.
