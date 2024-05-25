package services

import (
	"TCorp/grpc/models"
	"TCorp/proto/protocorp"
	"context"
)

type userGrpcServer struct {
	db models.DBUser
}

func (u userGrpcServer) SearchUser(ctx context.Context, req *protocorp.SearchUserRequest) (*protocorp.SearchUserResponse, error) {
	resp := new(protocorp.SearchUserResponse)
	filters := getFilters(req.Filter)
	response, err := u.db.SearchUser(filters)
	if err != nil {
		return nil, err
	}
	for _, v := range response {
		resp.Data = append(resp.Data, v.ToProto())
	}
	resp.Total = int64(len(response))
	return resp, nil
}

func getFilters(filter []*protocorp.SearchFilter) models.UserFilter {
	var m models.UserFilter
	for _, v := range filter {
		switch v.Key {
		case "name":
			m.Fname = v.Value
		case "city":
			m.City = v.Value
		case "phone":
			m.Phone = v.Value

		}
	}
	return m
}

func NewUserService(db models.DBUser) protocorp.UserInfoServer {
	return &userGrpcServer{db: db}
}

//gRPC controller to recieve the user input and calls service layer to get the user data based on requested user ID

func (u userGrpcServer) GetUserById(c context.Context, in *protocorp.GetUserByIdRequest) (*protocorp.GetUserByIdRespnose, error) {
	inputId := in.Id
	resp, err := u.db.GetUserById(int(inputId))
	if err != nil {
		return nil, err
	}
	var output protocorp.GetUserByIdRespnose
	var user protocorp.UserModel
	user.Id = int32(resp.ID)
	user.Name = string(resp.Fname)
	user.City = string(resp.City)
	user.Phone = resp.Phone
	user.Height = float32(resp.Height)
	user.Married = bool(resp.Married)
	output.User = &user
	return &output, nil
}

//gRPC controller to recieve the user input and calls service layer to get the users data based on requested user IDs

func (u userGrpcServer) GetUserByIds(c context.Context, in *protocorp.GetUserByIdsRequest) (*protocorp.GetUserByIdsResponse, error) {

	var inputId []int
	var userData []*protocorp.UserModel
	for i := range in.Id {
		inputId = append(inputId, int(in.Id[i]))
	}
	resp, err := u.db.GetUserByIds(inputId)
	if err != nil {
		return nil, err
	}
	var output protocorp.GetUserByIdsResponse

	for i := range resp {
		var user protocorp.UserModel
		user.Id = int32(resp[i].ID)
		user.Name = string(resp[i].Fname)
		user.City = string(resp[i].City)
		user.Phone = resp[i].Phone
		user.Height = float32(resp[i].Height)
		user.Married = bool(resp[i].Married)
		userData = append(userData, &user)
	}

	output.User = userData
	return &output, nil
}
