package db

import (
	"fmt"

	"github.com/alexeysamorodov/creator-applications/internal/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func ConnectDB(database config.Database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		database.Host,
		database.Port,
		database.User,
		database.Password,
		database.Name,
		database.SslMode,
	)

	db, err := sqlx.Open(database.Driver, dsn)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create database connection")

		return nil, err
	}

	return db, nil
}
