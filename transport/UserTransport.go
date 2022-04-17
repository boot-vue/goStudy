package transport

import (
	"context"
	"encoding/json"
	"errors"
	"goStudy/endpoint"
	"net/http"
	"strconv"
)

func HandleUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	if r.URL.Query().Get("id") != "" {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))

		return endpoint.UserRequest{
			UserId: id,
		}, nil
	}

	return nil, errors.New("参数错误")
}

func HandleUserResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
