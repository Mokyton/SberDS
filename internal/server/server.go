package server

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Host         string        `default:"localhost" envconfig:"HOST"`
	Port         string        `default:"8080" envconfig:"PORT"`
	WriteTimeout time.Duration `default:"10s" envconfig:"WRITE_TIMEOUT"`
	ReadTimeout  time.Duration `default:"10s" envconfig:"READ_TIMEOUT"`
}

func NewServer(cfg *Config, handler http.Handler) *http.Server {

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
}
