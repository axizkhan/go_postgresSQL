package models

import "time"

type UserResponse struct{
	ID int32 `json:"id"`
	Name string `json:"name"`
	DOB time.Time `json:"dob"`
	Age int `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}