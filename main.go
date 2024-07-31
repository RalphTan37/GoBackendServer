package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv" //package that grabs environment variables from a .env file

	//cmd line go mod vendor to copy code in vendor folder (local copy)
	"github.com/go-chi/chi"  //the chi router - third party router built on the same way the standard library in go does http routers
	"github.com/go-chi/cors" //cors (cross-origin resource sharing) configuration
)

func main() {
	godotenv.Load(".env") //loads .env file

	portString := os.Getenv("PORT") //reads the PORT var
	if portString == "" {
		log.Fatal("PORT is not found in the environment") //exits the program
	}

	router := chi.NewRouter() //creates new router object

	//can make request to the server from a browser
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//hooks up the handler to a specific HTTP method & path
	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness) //only on get requests
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)

	//server that the router can connect to
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on Port %v", portString)

	err := srv.ListenAndServe() //will block, code stops here and handles http requests

	//if anything goes wrong when handling requests
	if err != nil {
		log.Fatal(err)
	}
}
