package main

import (
	"SberDS/internal/config"
	"SberDS/internal/controllers"
	"SberDS/internal/server"
	"SberDS/internal/store"
	"log"
	"net/http"
)

type application struct {
	config *config.Config
	store  *store.Repository
	server *http.Server
}

func main() {
	var err error
	app := &application{}

	app.config, err = config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	app.store, err = store.New(app.config.StoreConfig)
	if err != nil {
		log.Fatal(err)
	}

	routes := struct {
		*controllers.ReportController
	}{controllers.NewReportController(app.store)}

	app.server = server.NewServer(app.config.ServerConfig, server.NewRouters(routes))

	app.server.ListenAndServe()
}
