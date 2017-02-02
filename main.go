package main

import (
	"log"

	"github.com/javinc/mango/database"
)

func main() {
	log.Println("[main]", "starting...")

	s := database.GetSession()

	log.Println(s)

	log.Println("[main]", "started!")
}
