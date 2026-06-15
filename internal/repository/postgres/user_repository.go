package postgres

import (
	"context"

	"github.com/axizkhan/go_postgresSQL/db/sqlc"
	appErrors "github.com/axizkhan/go_postgresSQL/internal/error"
	"github.com/axizkhan/go_postgresSQL/internal/repository"
	"github.com/jackc/pgx/v5"
)

type userRepository struct{
	queries *sqlc.Queries
}

func NewUserRepository(conn *pgx.Conn) repository.UserRepository{
	return &userRepository{
		queries: sqlc.New(conn),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, params sqlc.CreateUserParams)(sqlc.User, error){
	user,err:=r.queries.CreateUser(ctx,params)

	if err != nil{
		return  sqlc.User{},err
	}

	return  user, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int32)(sqlc.User, error){
	user,err:=r.queries.GetUserById(ctx, id)

	if err != nil{
		return  sqlc.User{},appErrors.ErrNotFound
	}

	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, params sqlc.UpdateUserParams)(sqlc.User, error){
	user,err := r.queries.UpdateUser(ctx,params);

	if err != nil{
		return  sqlc.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int32)error{
	return r.queries.DeleteUser(ctx,id)
}

func(r *userRepository) ListUsers(ctx context.Context, limit int32, offset int32)([]sqlc.User,error){
	users,err:=r.queries.ListUsers(ctx,sqlc.ListUsersParams{Limit: limit,Offset: offset})

	if err!=nil{
		return nil,err
	}

	return users,nil
}