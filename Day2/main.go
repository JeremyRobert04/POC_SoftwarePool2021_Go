package main

import (
	"SofwareGoDay2/server"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("env.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Port := os.Getenv("PORT")
	i := server.NewServer()
	i.Run(":" + Port)
}