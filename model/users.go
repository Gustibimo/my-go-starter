package model

import (
)

type User struct {
	ID int64 `bun:"id,pk,autoincrement"`
	Name string
	Email string
}