package models

import (
	"errors"
	"testing"
)

// Function for Testing GetUserById
func TestGetUserById(t *testing.T) {
	type testCase struct {
		input  int
		output User
		err    error
	}

	var cases []testCase = []testCase{
		{
			input: 1,
			output: User{

				ID:      1,
				Fname:   "Steve",
				City:    "LA",
				Phone:   "1234567890",
				Height:  5.8,
				Married: true,
			},
			err: nil,
		},
		{
			input:  20,
			output: User{},
			err:    errors.New("not found"),
		},
		{
			input:  -1,
			output: User{},
			err:    errors.New("invalid Id"),
		},
	}
	service := MockdbUser{}
	for i := range cases {
		resp, err := service.GetUserById(cases[i].input)
		if err != nil && err.Error() != cases[i].err.Error() {
			t.Error("failed test case error", i+1)
		}
		if resp.ID != cases[i].output.ID {
			t.Error("failed test case in output", i+1)
		}
	}

}

// Function for Testing GetUserByIds
func TestGetUserByIds(t *testing.T) {
	type testCase struct {
		input  []int
		output []User
		err    error
	}
	var cases []testCase = []testCase{

		{
			input: []int{
				1, 2,
			},
			output: []User{
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
			},
			err: nil,
		},

		{
			input: []int{
				25, 26, 27,
			},
			output: nil,
			err:    errors.New("not found"),
		},
		{
			input:  []int{},
			output: nil,
			err:    errors.New("id Cannot be empty"),
		},
	}
	u := MockdbUser{}
	for i := range cases {
		resp, err := u.GetUserByIds(cases[i].input)
		if err != nil && err.Error() != cases[i].err.Error() {
			t.Error("failed test case error", i+1)
		}

		if len(resp) != len(cases[i].output) {
			t.Error("failed test case in output", i+1)
		}
	}

}

func TestSearchFilter(t *testing.T) {
	type testCase struct {
		input  UserFilter
		output []User
		err    error
	}
	var cases []testCase = []testCase{

		{
			input: UserFilter{
				Fname: "Ganesh",
			},
			output: []User{
				{

					ID:      5,
					Fname:   "Ganesh",
					City:    "SA",
					Phone:   "4587890233",
					Height:  6.2,
					Married: true,
				},
			},
			err: nil,
		},

		{
			input: UserFilter{
				Fname: "",
				Phone: "",
				City:  "",
			},
			output: Users,
			err:    nil,
		},
		{
			input: UserFilter{
				Fname: "GANESH",
			},
			output: nil,
			err:    errors.New("not found"),
		},
	}
	u := MockdbUser{}
	for i := range cases {
		resp, err := u.SearchUser(cases[i].input)
		if err != nil && err.Error() != cases[i].err.Error() {
			t.Error("failed test case error", i+1)
		}
		if len(resp) != len(cases[i].output) {
			t.Error("failed test case in output", i+1)
		}
	}

}
