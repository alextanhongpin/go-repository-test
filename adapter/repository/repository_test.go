package repository_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "embed"

	"github.com/alextanhongpin/go-core-microservice/containers"
)

const postgresVersion = "15.1-alpine"

func TestMain(m *testing.M) {
	stop := containers.StartPostgres(postgresVersion, func(db *sql.DB) error {
		b, err := os.ReadFile("../postgres/schemas/schema.sql")
		if err != nil {
			return err
		}

		_, err = db.Exec(string(b))
		if err != nil {
			return err
		}

		log.Println("database migration completed")

		return nil
	})
	code := m.Run()
	stop()
	os.Exit(code)
}
