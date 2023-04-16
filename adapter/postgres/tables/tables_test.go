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

//go:embed testdata/baseline.sql
var baseline string

func TestMain(m *testing.M) {
	stop := containers.StartPostgres(postgresVersion, func(db *sql.DB) error {
		// Issue, there is no easy way to run the migration.
		_, err := db.Exec(baseline)
		if err != nil {
			return err
		}

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