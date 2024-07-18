package main

import "github.com/gen2brain/beeep";

// Luo ilmoituksen, jossa otsikko on "Pörssisähkö", viesti on funktiolle annettu parametri (message) ja kuvake on sahko.png
func createNotification(message string) {
	beeep.Notify("Pörssisähkö", message, "./sahko.png")
}

