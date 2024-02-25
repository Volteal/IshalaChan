package main

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func initialize() (*discordgo.Session, string, error) {
	godotenv.Load("./env/.env")

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("You forogt to set your api_key.")
	}

	prefix := os.Getenv("PREFIX")
	if prefix == "" {
		log.Fatal("You forogt to set your prefix.")
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	return s, prefix, nil
}
