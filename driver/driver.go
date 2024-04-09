// driver/driver.go
package driver

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

// ConnectSQL creates a connection pool for PostgreSQL using pgx v5.
func ConnectSQL(dsn string) (*DB, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	// Set pool settings from environment variables
	maxOpenDbConn, err := strconv.Atoi(os.Getenv("MAX_OPEN_DB_CONN"))
	if err != nil {
		return nil, err
	}

	maxIdleDbConn, err := strconv.Atoi(os.Getenv("MAX_IDLE_DB_CONN"))
	if err != nil {
		return nil, err
	}

	maxDbLifetimeMinutes, err := strconv.Atoi(os.Getenv("MAX_DB_LIFETIME_MINUTES"))
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(maxOpenDbConn)
	config.MinConns = int32(maxIdleDbConn)

	config.MaxConnLifetime = time.Duration(maxDbLifetimeMinutes) * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	db := &DB{Pool: pool}
	return db, nil
}

func (db *DB) PrintPoolStats() {
	log.Printf("Total connections (from pool stats): %v", db.Pool.Stat().TotalConns())
	log.Printf("Idle connections (from pool stats): %v", db.Pool.Stat().IdleConns())
	log.Printf("Acquired Duration (from pool stats): %v", db.Pool.Stat().AcquireDuration())
}
