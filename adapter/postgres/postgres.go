package postgres

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/amacneil/dbmate/pkg/dbmate"
	_ "github.com/amacneil/dbmate/pkg/driver/postgres"
)

// Migrate executes the migration programmatically.
// Actually, we can just run the migration in schemas/schema.sql
func Migrate(ds string, path string) error {
	u, err := url.Parse(ds)
	if err != nil {
		return err
	}
	migrator := dbmate.New(u)

	// Always point the path to the root go.mod file relative to current directory.
	// e.g. Migrate(ds, "../../go.mod")
	dir := filepath.Dir(path)
	migrator.MigrationsDir = filepath.Join(dir, "./adapter/postgres/migrations/")
	migrator.SchemaFile = filepath.Join(dir, "./adapter/postgres/schemas/schema.sql")
	err = migrator.CreateAndMigrate()
	if err != nil {
		return fmt.Errorf("failed to migrate: %s", err)
	}

	return nil
}
