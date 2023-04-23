package tables

import (
	"context"

	"github.com/uptrace/bun"
)

// client represents either a db or tx
// client.
type client interface {
	DB(ctx context.Context) bun.IDB
}

type OffsetPagination struct {
	Limit   int
	Offset  int
	OrderBy []string
}

type Relation struct {
	Name  string
	Apply []func(*bun.SelectQuery) *bun.SelectQuery
}
