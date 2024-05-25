package controller

import (
	"TCorp/proto/protocorp"
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controller struct {
	Service protocorp.UserInfoClient
}
type UserController interface {
	SearchUser(c echo.Context) error
	GetUserById(c echo.Context) error
	GetUserByIds(c echo.Context) error
}

func NewController(service protocorp.UserInfoClient) UserController {
	return &controller{Service: service}
}

// function to search user based on filter
func (c2 controller) SearchUser(c echo.Context) error {
	filter := getFilters(c.QueryParams())

	resp, err := c2.Service.SearchUser(context.Background(), &protocorp.SearchUserRequest{Filter: filter})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

// function to get specific user data
func (c2 controller) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := c2.Service.GetUserById(context.Background(), &protocorp.GetUserByIdRequest{Id: int32(id)})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

// function to get user data based on IDs provided
func (c2 controller) GetUserByIds(c echo.Context) error {
	type Payload struct {
		IDs []int32 `json:"ids"`
	}
	var payload Payload
	if err := c.Bind(&payload); err != nil {
		return err
	}
	// todo get the filter and pass it to search
	resp, err := c2.Service.GetUserByIds(context.Background(), &protocorp.GetUserByIdsRequest{Id: payload.IDs})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, resp)
}

// helper function to get the input filters
func getFilters(f url.Values) []*protocorp.SearchFilter {
	var filters []*protocorp.SearchFilter
	for k, v := range f {
		switch k {
		case "name":
			filters = append(filters, &protocorp.SearchFilter{
				Key:   k,
				Value: v[0],
			})
		case "phone":
			filters = append(filters, &protocorp.SearchFilter{
				Key:   k,
				Value: v[0],
			})
		case "city":
			filters = append(filters, &protocorp.SearchFilter{
				Key:   k,
				Value: v[0],
			})

		}
	}
	return filters
}
