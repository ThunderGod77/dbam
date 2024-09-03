package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/ThunderGod77/dbam/internal/core"

	_ "github.com/lib/pq"
)

type postgresRepo struct {
	sync.RWMutex

	core.ConnObject
	db *sql.DB
}

func NewPostgresService(conn core.ConnObject) (core.DbDataService, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conn.Host, conn.Port, conn.User, conn.Password, conn.DbName, conn.SslMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &postgresRepo{
		RWMutex:    sync.RWMutex{},
		ConnObject: conn,
		db:         db,
	}, nil
}
