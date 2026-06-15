package models

import "time"

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	DOB time.Time `json:"dob" validate:"required"`
}

type ListUserRequest struct {
	Offeset int32 `json:"offeset" validate:"required"`
	Limit int32 `json:"limit" validate:"required"`
}

type GetUserRequest struct {
	ID int32 `json:"id" validate:"required"`
}

type DeleteUserRequest struct {
	ID int32 `json:"id" validate:"required"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	DOB time.Time `json:"dob" validate:"required"`
}