package store

import (
	"SberDS/internal/models"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type Config struct {
	DBHost   string `default:"localhost" envconfig:"HOST"`
	DBPort   int    `default:"5432" envconfig:"PORT"`
	DBName   string `default:"sandbox" envconfig:"NAME"`
	Username string `default:"postgres" envconfig:"USERNAME"`
	Password string `default:"postgres123" envconfig:"PASSWORD"`
	LifeTime string `default:"10m" envconfig:"LIFE_TIME"`
	PoolSize int    `default:"5" envconfig:"POOLSIZE"`
	MaxIdle  int    `default:"3" envconfig:"MAXIDLE"`
}

type Repository struct {
	db     *sql.DB
	config *Config
}

func New(cfg *Config) (*Repository, error) {
	src := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.Username, cfg.Password)

	db, err := sql.Open("postgres", src)
	if err != nil {
		return nil, err
	}

	lifeDuration, err := time.ParseDuration(cfg.LifeTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(lifeDuration)
	db.SetMaxOpenConns(cfg.PoolSize)
	db.SetMaxIdleConns(cfg.MaxIdle)

	return &Repository{
		db:     db,
		config: cfg,
	}, nil

}

func (r *Repository) GetReportById(ctx context.Context, reportId int) (string, error) {
	var res string

	if err := r.db.QueryRowContext(ctx, getReportByIDSQL, reportId).Scan(&res); err != nil {
		return "", fmt.Errorf("can't get report: %v", err)
	}

	return res, nil
}

func (r *Repository) SetReport(ctx context.Context, data models.SetReportReq) error {

	_, err := r.db.ExecContext(ctx, addReportSQL, data.ReportInfo, data.ModelID)
	if err != nil {
		return fmt.Errorf("can't set report: %v", err)
	}

	return nil
}
