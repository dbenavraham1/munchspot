package main

import (
	"github.com/dbenavraham1/munchspot/app"
	"github.com/dbenavraham1/munchspot/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

func main() {
	appPath, err := app.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	newApp, err := app.NewApp(appPath)
	if err != nil {
		log.Fatal(err)
	}

	rtr := mux.NewRouter()
	rtr.Path("/geolocation/{format}").HandlerFunc(controllers.GeocodeLocationHandler).Methods("GET")
	rtr.Path("/food/resource/{id}/{format}").HandlerFunc(controllers.FoodResourceHandler).Methods("GET")
	handler := cors.AllowAll().Handler(rtr)

	server := &http.Server {
		Addr:         newApp.Server.Host + ":" + newApp.Server.Port,
		Handler:      handler,
		ReadTimeout:  newApp.Server.Timeout.Read * time.Second,
		WriteTimeout: newApp.Server.Timeout.Write * time.Second,
		IdleTimeout:  newApp.Server.Timeout.Idle * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
