package transport

import (
	"github.com/gophers0/users/internal/model"
	"github.com/gophers0/users/pkg/errs"
)

type CreateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (req *CreateUserRequest) Validate() error {
	if req.Login == "" {
		return errs.RequestValidationError.AddInfo("empty login")
	}
	return nil
}

type CreateUserResponse struct {
	BaseResponse
	User *model.User `json:"user"`
}

type UpdateUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (req *UpdateUserRequest) Validate() error {
	if req.Login == "" {
		return errs.RequestValidationError.AddInfo("empty login")
	}
	return nil
}

type UpdateUserResponse struct {
	BaseResponse
	User *model.User `json:"user"`
}
