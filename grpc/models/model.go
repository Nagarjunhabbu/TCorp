package models

import "TCorp/proto/protocorp"

type User struct {
	ID      int     `json:"id"`
	Fname   string  `json:"fname"`
	City    string  `json:"city"`
	Phone   string  `json:"phone"`
	Height  float64 `json:"height"`
	Married bool    `json:"Married"`
}

type UserFilter struct {
	Fname string
	City  string
	Phone string
}

// User Database
var Users []User = []User{
	{

		ID:      1,
		Fname:   "Steve",
		City:    "LA",
		Phone:   "1234567890",
		Height:  5.8,
		Married: true,
	},
	{

		ID:      2,
		Fname:   "Virat",
		City:    "CA",
		Phone:   "4567890233",
		Height:  6.2,
		Married: true,
	},

	{

		ID:      3,
		Fname:   "Alex",
		City:    "AL",
		Phone:   "8907654323",
		Height:  6.1,
		Married: false,
	},
	{

		ID:      4,
		Fname:   "Rohit",
		City:    "CA",
		Phone:   "4567890233",
		Height:  6.3,
		Married: false,
	},
	{

		ID:      5,
		Fname:   "Ganesh",
		City:    "SA",
		Phone:   "4587890233",
		Height:  6.2,
		Married: true,
	},
}

func (u User) ToProto() *protocorp.UserModel {
	return &protocorp.UserModel{
		Id:      int32(u.ID),
		Name:    u.Fname,
		City:    u.City,
		Phone:   u.Phone,
		Height:  float32(u.Height),
		Married: u.Married,
	}
}
