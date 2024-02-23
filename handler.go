package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Volteal/IshalaChan/commands"
	"github.com/bwmarrin/discordgo"
)

func handler(sess *discordgo.Session, prefix string) {
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		//Server Logic
		args := strings.Split(m.Content, " ")
		if args[0] != prefix {
			return
		}

		if args[1] == "hello" {
			commands.HelloWorldHandler(s, m)
		}

		if args[1] == "quote" {
			commands.QuotesHandler(s, m)
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err := sess.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer sess.Close()

	fmt.Println("IshalaChan is Online!")
}
