package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" //package that grabs environment variables from a .env file
	//cmd line go mod vendor to copy code in vendor folder (local copy)
)

func main() {
	godotenv.Load(".env") //loads .env file

	portString := os.Getenv("PORT") //reads the PORT var
	if portString == "" {
		log.Fatal("PORT is not found in the environment") //exits the program
	}

	fmt.Println("Port:", portString)
}
