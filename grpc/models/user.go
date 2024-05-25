package models

import (
	"errors"
)

type dbUser struct {
}

type DBUser interface {
	GetUserById(Id int) (User, error)
	GetUserByIds(Ids []int) ([]User, error)
	SearchUser(filter UserFilter) ([]User, error)
}

func NewUserDb() DBUser {
	return dbUser{}
}

// GetUserById returns the user data based on UserId
func (u dbUser) GetUserById(Id int) (User, error) {
	if Id <= 0 {
		return User{}, errors.New("invalid Id")
	}

	for _, v := range Users {
		if v.ID == Id {
			return v, nil
		}
	}
	return User{}, errors.New("not found")
}

// GetUserById returns the users data based on UserIds
func (u dbUser) GetUserByIds(Ids []int) ([]User, error) {
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

// searchUser return the users data based on filter passsed
func (d dbUser) SearchUser(filter UserFilter) ([]User, error) {
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
