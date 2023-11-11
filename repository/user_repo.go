package repository

import (
	"context"
	"my-go-starter/db"
	"my-go-starter/model"
	"my-go-starter/shared"
)

type UserRepo interface {
	Save(model.User) error
	Get(name string) (*model.User, error)
	FindAll() ([]model.User, error)
}


func FindAll() ([]model.User, error) {
	var users []model.User

	users = append(users, model.User{
		ID: 1,
		Name: "John",
		Email: "jo@marlaw.id",
	})

	users = append(users, model.User{
		ID: 2,
		Name: "Doe",
		Email: "doe@marlaw.id",
	})

	// err := db.Bun.NewSelect().Model(&users).Scan(context.Background())
	return users, nil
}

func Save(user shared.UserType) (*model.User, error) {

	u := &model.User{
		Name: user.Name,
		Email: user.Email,
	}

	_, err := db.Bun.NewInsert().Model(user).Exec(context.Background())
	return u, err
}