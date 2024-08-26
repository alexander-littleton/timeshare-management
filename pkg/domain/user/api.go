package user

import (
	"context"
	"encoding/json"
	"net/http"
)

type UserService interface {
	CreateUser(ctx context.Context, user User) error
}

type Controller struct {
	userService UserService
}

func NewController(userService UserService) Controller {
	return Controller{
		userService: userService,
	}
}

func (c Controller) CreateUser(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.userService.CreateUser(ctx, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
