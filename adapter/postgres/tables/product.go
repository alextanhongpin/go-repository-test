package tables

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID          int64 `bun:",pk"`
	Name        string
	Description string
	UserID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Relations.
	User *User `bun:"rel:belongs-to,join:user_id=id"`
}

type ProductTable struct {
	client
}

func NewProduct(client client) *ProductTable {
	return &ProductTable{
		client: client,
	}
}

func (t *ProductTable) Find(ctx context.Context, id int64, relations ...Relation) (*Product, error) {
	p := &Product{
		ID: id,
	}

	q := t.DBTx(ctx).NewSelect().Model(p)
	for _, relation := range relations {
		q = q.Relation(relation.Name, relation.Apply...)
	}

	if err := q.WherePK().Scan(ctx); err != nil {
		return nil, err
	}

	return p, nil
}

func (t *ProductTable) List(ctx context.Context, pagination *OffsetPagination) ([]Product, error) {
	var p []Product

	q := t.DBTx(ctx).NewSelect().Model(&p)

	if pagination != nil {
		q = q.
			Limit(pagination.Limit).
			Offset(pagination.Offset)

		if len(pagination.OrderBy) > 0 {
			q = q.Order(pagination.OrderBy...)
		}
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return p, nil
}

func (t *ProductTable) Create(ctx context.Context, name, description string, userID uuid.UUID) (*Product, error) {
	p := &Product{
		Name:        name,
		Description: description,
		UserID:      userID,
	}

	if _, err := t.DBTx(ctx).
		NewInsert().
		Model(p).
		Column(
			"name",
			"description",
			"user_id",
		).
		Returning("*").
		Exec(ctx); err != nil {
		return nil, err
	}

	return p, nil
}

func (t *ProductTable) Delete(ctx context.Context, id int64) error {
	p := &Product{
		ID: id,
	}

	_, err := t.DBTx(ctx).
		NewDelete().
		Model(p).
		WherePK().
		Exec(ctx)

	return err
}
