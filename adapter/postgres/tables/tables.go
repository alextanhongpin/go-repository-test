package tables

import (
	"context"

	"github.com/alextanhongpin/dbtx/buntx"
	"github.com/uptrace/bun"
)

// client represents either a db or tx
// client.
type client interface {
	DBTx(ctx context.Context) buntx.DBTX
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
