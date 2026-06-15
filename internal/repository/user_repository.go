package repository

import (
	"context"

	"github.com/axizkhan/go_postgresSQL/db/sqlc"
)

type UserRepository interface{
	CreateUser(ctx context.Context, params sqlc.CreateUserParams)(sqlc.User,error)

	GetUserById(ctx context.Context,id int32)(sqlc.User,error)

	UpdateUser(ctx context.Context,params sqlc.UpdateUserParams)(sqlc.User,error)

	DeleteUser(ctx context.Context, id int32)error

	ListUsers(ctx context.Context,limit int32, offset int32)([]sqlc.User,error)
}