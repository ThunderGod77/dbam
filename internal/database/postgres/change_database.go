package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

func (pr *postgresRepo) ChangeDb(dbName string) error {
	pr.Lock()
	defer pr.Unlock()

	pr.DbName = dbName

	err := pr.db.Close()
	if err != nil {
		return err
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pr.Host, pr.Port, pr.User, pr.Password, pr.DbName, pr.SslMode,
	)

	pr.db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	log.Println("successfully changed the database")

	return nil
}
