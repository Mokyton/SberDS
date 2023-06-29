package main

import (
	"SberDS/internal/config"
	"SberDS/internal/controllers"
	"SberDS/internal/server"
	"SberDS/internal/store"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	b := controllers.NewBaseController(app.store)

	routes := struct {
		*controllers.MiddlewareController
		*controllers.ReportController
	}{
		controllers.NewMiddlewareController(b),
		controllers.NewReportController(b),
	}

	app.server = server.NewServer(app.config.ServerConfig, server.NewRouters(routes))

	serverErrors := make(chan error, 1)
	go func() {
		log.Println("Listen and serve on:", fmt.Sprintf("%s:%s", app.config.ServerConfig.Host, app.config.ServerConfig.Port))
		serverErrors <- app.server.ListenAndServe()
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("Server start error: %v", err)
	case <-osSignals:
		log.Println("Shutdown...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := app.server.Shutdown(ctx); err != nil {
			log.Printf("Gracefull shutdown in %v not completed, error: %v", time.Second*5, err)
			if err := app.server.Close(); err != nil {
				log.Fatalf("Server stop error: %v", err)
			}
		}
	}

}
