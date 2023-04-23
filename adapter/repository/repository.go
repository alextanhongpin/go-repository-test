package repository

import (
	"context"

	"github.com/uptrace/bun"
)

type atomic interface {
	DB(ctx context.Context) bun.IDB
	RunInTx(ctx context.Context, fn func(context.Context) error) error
}

func Map[K, V any](ks []K, fn func(i int) V) []V {
	res := make([]V, len(ks))
	for i := 0; i < len(ks); i++ {
		res[i] = fn(i)
	}

	return res
}
