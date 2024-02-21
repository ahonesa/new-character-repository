package main

import (
	"flag"
	"github.com/ahonesa/new-character-repository/internal/config"
	"github.com/ahonesa/new-character-repository/internal/cthulhu"
	"github.com/ahonesa/new-character-repository/internal/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	envFile := flag.String("env", "local.env", "Specify the environment file to use.")
	flag.Parse()

	// Load the application configuration.
	cfg := config.LoadConfig(*envFile)

	// Connect to MongoDB.
	db.ConnectToMongoDB(cfg)

	r := mux.NewRouter()

	// Set up Cthulhu routes
	r.HandleFunc("/api/cthulhu/characters/{id}", cthulhu.GetCthulhuCharacterByID).Methods("GET")
	r.HandleFunc("/api/cthulhu/characters", cthulhu.GetAllCthulhuCharacters).Methods("GET")
	r.HandleFunc("/api/cthulhu/characters", cthulhu.CreateCharacterHandler).Methods("POST")
	r.HandleFunc("/api/cthulhu/characters/{id}", cthulhu.UpdateCharacterHandler).Methods("PUT")
	r.HandleFunc("/api/cthulhu/characters/{id}", cthulhu.DeleteCharacterHandler).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))

}
