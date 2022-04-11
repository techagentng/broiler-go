package main

import (
	"github.com/techagentng/boiler-go/application/server"
	"github.com/techagentng/boiler-go/domain/helpers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	helpers.InitializeLogDir()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env with godotenv: %s", err)
	}
	server.Start()
}
