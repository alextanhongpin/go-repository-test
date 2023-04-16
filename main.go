package main

import (
	"context"
	"fmt"

	"github.com/alextanhongpin/go-core-microservice/database/postgres"
	"github.com/alextanhongpin/go-repository-test/adapter/repository"
	"github.com/alextanhongpin/uow/bun"
)

func main() {
	db := postgres.NewBun()
	uow := bun.New(db)

	authRepo := repository.NewAuth(uow)
	ctx := context.Background()
	user, err := authRepo.CreateUser(ctx, "john appleseed")
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created: %#v\n", user)

	userID := user.ID
	{
		user, err := authRepo.FindUserByID(ctx, userID)
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
