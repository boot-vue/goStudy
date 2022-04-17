package main

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"goStudy/endpoint"
	"goStudy/service"
	"goStudy/transport"
	"net/http"
)

func main() {
	userService := service.UserService{}
	uerEndpoint := endpoint.UserEndpoint(userService)

	s := kithttp.NewServer(uerEndpoint, transport.HandleUserRequest, transport.HandleUserResponse)

	http.ListenAndServe(":8080", s)
}
