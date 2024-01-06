package testfixture

import (
	"context"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/google/uuid"
)

func CreateProducts(t *testing.T, ctx context.Context, c client, n int, variants ...string) []tables.Product {
	userID := CreateUser(t, ctx, c).ID
	products := make([]tables.Product, n)
	for i := 0; i < n; i++ {
		products[i] = *CreateProduct(t, ctx, c, append(variants, "with_user:"+userID.String())...)
	}

	return products
}

func CreateProduct(t *testing.T, ctx context.Context, c client, variants ...string) *tables.Product {
	t.Helper()

	p := NewProduct(append(variants, "db:create")...)
	if p.UserID == uuid.Nil {
		p.UserID = CreateUser(t, ctx, c).ID
	}

	tbl := tables.NewProduct(c)
	product, err := tbl.Create(ctx, p.Name, p.Description, p.UserID)
	if err != nil {
		t.Fatalf("failed to create product: %v", err)
	}

	store(t.Name(), "products", product.ID, product)

	return product
}

func NewProduct(variants ...string) *tables.Product {
	p := &tables.Product{
		ID:          1,
		Name:        "table",
		Description: "a wooden table",
		UserID:      uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	for _, v := range variants {
		if strings.HasPrefix(v, "with_user:") {
			_, userID, _ := strings.Cut(v, "with_user:")
			p.UserID = uuid.MustParse(userID)
			continue
		}

		switch v {
		case "db:create":
			p.ID = 0
			p.UserID = uuid.Nil
		default:
			log.Fatalf("unknown product variant: %s", v)
		}
	}

	return p
}
