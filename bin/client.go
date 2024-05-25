package main

import (
	"TCorp/proto/protocorp"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("connection failed")
	}
	defer conn.Close()
	serv := protocorp.NewUserInfoClient(conn)
	// getuserbyId
	resp1, _ := serv.GetUserById(context.Background(), &protocorp.GetUserByIdRequest{
		Id: 1,
	})
	fmt.Println(resp1)

	//getUserByIds
	resp, _ := serv.GetUserByIds(context.Background(), &protocorp.GetUserByIdsRequest{
		Id: []int32{
			3, 4,
		},
	})
	fmt.Println(resp)

	//SearchUser
	resp2, _ := serv.SearchUser(context.Background(), &protocorp.SearchUserRequest{
		Filter: []*protocorp.SearchFilter{
			{Key: "name", Value: "Ganesh"},
		},
	})
	fmt.Println(resp2)
}
