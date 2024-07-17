package main

import "github.com/gen2brain/beeep";

func createNotification(message string) {
	beeep.Notify("Pörssisähkö", message, "./sahko.png")
}

