package postgres

import (
	"embed"
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/alextanhongpin/go-repository-test/internal"
	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
)

//go:embed migrations/*.sql
var migrations embed.FS

// Migrate executes the migration programmatically.
// Actually, we can just run the migration in schemas/schema.sql
func Migrate(dsn string) error {
	u, err := url.Parse(dsn)
	if err != nil {
		return err
	}
	db := dbmate.New(u)
	db.FS = migrations

	db.MigrationsDir = []string{"migrations/"}
	db.SchemaFile = filepath.Join(internal.Root, "./adapter/postgres/schemas/schema.sql")
	err = db.CreateAndMigrate()
	if err != nil {
		return fmt.Errorf("failed to migrate: %s", err)
	}

	return nil
}
