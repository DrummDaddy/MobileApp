package main

import (
	"MobileApp/database"
	"MobileApp/routes"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	http.HandleFunc("/register", routes.RegisterHendler(db))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
