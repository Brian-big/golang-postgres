package main

import (
	"brian-big/go-ms/api"
	"brian-big/go-ms/db"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Server has started!")

	pgdb, err := db.StartDB()

	if err != nil {
		log.Printf("Error starting the database: %v\n", err)
	}

	router := api.RunApi(pgdb)

	port := os.Getenv("PORT")

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)

	if err != nil {
		log.Printf("Error from router: %v\n", err)
	}

}
