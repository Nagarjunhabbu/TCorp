package main

import (
	"TCorp/http/controller"
	"TCorp/proto/protocorp"
	"context"
	"log"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ctx := context.Background()
	userServiceUrl := "grpcservice:5001"
	conn, err := grpc.DialContext(ctx, userServiceUrl, grpc.WithInsecure())
	if err != nil {
		log.Println("error in initializing the grpc")
	}

	serv := protocorp.NewUserInfoClient(conn)
	ctrl := controller.NewController(serv)

	e.GET("/v1/user", ctrl.SearchUser)
	e.GET("/v1/user/:id", ctrl.GetUserById)
	e.POST("/v1/user", ctrl.GetUserByIds)

	e.Logger.Fatal(e.Start(":8000"))
}
