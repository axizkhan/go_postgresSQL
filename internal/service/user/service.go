package user

import (
	"context"
	"time"

	"github.com/axizkhan/go_postgresSQL/db/sqlc"
	"github.com/axizkhan/go_postgresSQL/internal/models"
	"github.com/axizkhan/go_postgresSQL/internal/repository"
)

type Service struct{
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository,) *Service{
	return &Service{
		repo: repo,
	}
}

func calculateAge(dob time.Time) int{
	
	now:= time.Now()
	age:= now.Year()-dob.Year()

	if now.YearDay()<dob.YearDay(){
		age--
	}

	return age
}

func mapUserResponse(user sqlc.User) models.UserResponse{
	return models.UserResponse{
		ID: user.ID,
		Name: user.Name,
		DOB: user.Dob.Time,
		Age:calculateAge(user.Dob.Time),
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,

	}
}

func(s *Service) CreateUser(ctx context.Context, params sqlc.CreateUserParams)(models.UserResponse,error){
	user,err:=s.repo.CreateUser(ctx,params)
	if err !=nil{
		return models.UserResponse{},err
	}
	return mapUserResponse(user),nil
}

func(s *Service) GetUserById(ctx context.Context, id int32)(models.UserResponse,error){
	user,err:=s.repo.GetUserById(ctx,id)
	if err != nil{
		return  models.UserResponse{},err
	}

	return  mapUserResponse(user),nil
}

func(s *Service) ListUsers(ctx context.Context,limit int32, offset int32)([]models.UserResponse,error){
	users,err:=s.repo.ListUsers(ctx,limit,offset)

	if err!=nil{
		return nil,err
	}

	response:=make([]models.UserResponse,0)

	for _, user := range users {
		response = append(
			response,
			mapUserResponse(user),
		)
	}

	return response, nil
}

func(s *Service) UpdateUser(ctx context.Context, params sqlc.UpdateUserParams)(models.UserResponse,error){
	user,err:=s.repo.UpdateUser(ctx,params)

	if err != nil{
		return models.UserResponse{},err
	}

	return mapUserResponse(user),nil
}

func(s*Service) DeleteUser(ctx context.Context, id int32)error{
	return s.repo.DeleteUser(ctx,id)
}