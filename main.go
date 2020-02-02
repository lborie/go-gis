package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/lborie/go-gis/dao"
	"github.com/lborie/go-gis/handlers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RenderMap)
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	r.HandleFunc("/regions", handlers.Regions)
	r.HandleFunc("/departements", handlers.Departements)

	var serverPort = "80"
	if os.Getenv("APPSETTING_PORT") != "" {
		serverPort = os.Getenv("APPSETTING_PORT")
	}
	server := &http.Server{
		Addr:           ":" + serverPort,
		Handler:        r,
		ReadTimeout:    5 * time.Minute,
		WriteTimeout:   5 * time.Minute,
	}

	// Log Configuration
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		QuoteEmptyFields:       true,
		FullTimestamp:          true,
		ForceColors:            true,
		DisableLevelTruncation: true,
	})

	// Init databse connection
	connectionURI := os.Getenv("DB_CONNECTION_URI")
	db, err := dao.InitDatabasePostgreSQL(connectionURI)
	if err != nil {
		log.Errorf("Connection to database impossible : %s", err.Error())
	}
	dao.DB = db

	// Launch
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	// Graceful Shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Info("Got SIGINT...")
	case syscall.SIGTERM:
		log.Info("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	_ = server.Shutdown(context.Background())
	log.Print("Done")
}
