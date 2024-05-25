package models

import (
	"errors"
)

type MockdbUser struct {
}

// GetUserById returns the user data based on UserId

func (u MockdbUser) GetUserById(Id int) (User, error) {
	if Id <= 0 {
		return User{}, errors.New("invalid Id")
	}

	if Id == 1 {
		return User{
			ID:      1,
			Fname:   "Steve",
			City:    "LA",
			Phone:   "1234567890",
			Height:  5.8,
			Married: true,
		}, nil
	}
	return User{}, errors.New("not found")
}

// GetUserById returns the users data based on UserIds
func (u MockdbUser) GetUserByIds(Ids []int) ([]User, error) {
	if len(Ids) <= 0 {
		return nil, errors.New("id Cannot be empty")

	}
	var users []User
	for i := range Ids {
		for _, v := range Users {
			if v.ID == Ids[i] {
				users = append(users, v)
			}
		}
	}
	if len(users) > 0 {
		return users, nil
	}
	return nil, errors.New("not found")
}

func (d MockdbUser) SearchUser(filter UserFilter) ([]User, error) {
	var resp []User
	if filter.Fname == "" && filter.City == "" && filter.Phone == "" {
		return Users, nil
	}

	for _, v := range Users {
		if filter.Phone != "" && v.Phone == filter.Phone {
			resp = append(resp, v)
			continue
		}
		if filter.City != "" && v.City == filter.City {
			resp = append(resp, v)
			continue
		}
		if filter.Fname != "" && v.Fname == filter.Fname {
			resp = append(resp, v)
			continue
		}

	}
	if len(resp) <= 0 {
		return nil, errors.New("not found")
	}
	return resp, nil
}
