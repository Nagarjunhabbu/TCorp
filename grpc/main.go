package main

import (
	"TCorp/grpc/models"
	"TCorp/grpc/services"
	"TCorp/proto/protocorp"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Registering gRPC server
	s := grpc.NewServer()
	db := models.NewUserDb()

	server := services.NewUserService(db)
	protocorp.RegisterUserInfoServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
