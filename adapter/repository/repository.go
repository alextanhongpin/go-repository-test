package repository

import (
	"context"

	"github.com/alextanhongpin/dbtx/buntx"
)

type atomic interface {
	DBTx(ctx context.Context) buntx.DBTX
	RunInTx(ctx context.Context, fn func(context.Context) error) error
}

func Map[K, V any](ks []K, fn func(i int) V) []V {
	res := make([]V, len(ks))
	for i := 0; i < len(ks); i++ {
		res[i] = fn(i)
	}

	return res
}
