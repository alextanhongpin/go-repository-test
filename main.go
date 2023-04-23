package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alextanhongpin/dbtx/buntx"
	xpostgres "github.com/alextanhongpin/go-core-microservice/database/postgres"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres"
	"github.com/alextanhongpin/go-repository-test/adapter/repository"
)

func main() {
	dsn := postgres.NewDSN()
	if err := postgres.Migrate(dsn, "."); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	db := xpostgres.NewBun()
	btx := buntx.New(db)

	authRepo := repository.NewAuthRepository(btx)
	ctx := context.Background()
	user, err := authRepo.CreateUser(ctx, "john appleseed")
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created: %#v\n", user)

	userID := user.ID
	{
		user, err := authRepo.FindUser(ctx, userID)
		if err != nil {
			panic(err)
		}

		fmt.Printf("User queried: %#v\n", user)
	}

	if err := authRepo.Delete(ctx, userID); err != nil {
		panic(err)
	}
	fmt.Println("User deleted")
}
