package main

import (
	"log"

	"github.com/hararudoka/pkggobot/internal/bot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	b, err := bot.New("bot.yml")
	if err != nil {
		log.Fatal(err)
	}

	b.Start()
}