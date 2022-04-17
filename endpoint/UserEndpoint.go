package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"goStudy/service"
)

type UserRequest struct {
	UserId int `json:"userId"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func UserEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserRequest)
		result := userService.GetUser(req.UserId)
		return UserResponse{Result: result}, nil
	}
}
