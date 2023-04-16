package tables_test

import (
	"bytes"
	"database/sql"
	"log"
	"os"
	"os/exec"
	"testing"

	_ "embed"

	"github.com/alextanhongpin/go-core-microservice/containers"
)

const postgresVersion = "15.1-alpine"

func TestMain(m *testing.M) {
	stop := containers.StartPostgres(postgresVersion, func(db *sql.DB) error {
		b, err := os.ReadFile("../schemas/schema.sql")
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

func cmd(arg string, args ...string) {
	cmd := exec.Command(arg, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	log.Printf("sdtout: %s", stdout.String())
	log.Printf("sderr: %s", stderr.String())
}
