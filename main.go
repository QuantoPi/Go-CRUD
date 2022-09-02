package main

import (
	"log"
	"net/http"
	"time"

	"github.com/QuantoPi/Go-CRUD/db"
	"github.com/QuantoPi/Go-CRUD/models"
	"github.com/QuantoPi/Go-CRUD/routes"
	"github.com/gorilla/mux"
)

func main() {
	// db connection and migrations
	db.DBConnection()
	db.DB.AutoMigrate(models.Costumer{})

	//server configuration and routes
	r := mux.NewRouter()

	//css and js routes
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//home and html routes
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/db-handler", routes.DatabaseHandler)
	r.HandleFunc("/costumers-table", routes.CostumersHandler)
	//api routes
	r.HandleFunc("/costumers", routes.GetCostumersHandler).Methods("GET")
	r.HandleFunc("/costumers/{id}", routes.GetCostumerHandler).Methods("GET")
	r.HandleFunc("/costumers", routes.PostCostumerHandler).Methods("POST")
	r.HandleFunc("/costumers/{id}", routes.DeleteCostumerHandler).Methods("DELETE")
	r.HandleFunc("/costumers/{id}", routes.UpdateCostumerHandler).Methods("PUT")
	//ip and server config
	srv := &http.Server{
		Handler: r,
		// Dont open with 'https://'
		Addr: "127.0.0.1:8080",
		//check server timeout
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
