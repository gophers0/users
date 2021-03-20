package handlers

import (
	"net/http"
	"strconv"

	"github.com/gophers0/users/internal/transport"
	"github.com/gophers0/users/pkg/crypto"
	"github.com/gophers0/users/pkg/errs"
	"github.com/labstack/echo"
)

func (h *Handlers) CreateUser(c echo.Context) error {
	req := &transport.CreateUserRequest{}
	if err := transport.BindAndValidate(c, req); err != nil {
		return errs.NewStack(err)
	}

	cryptedPwd, err := crypto.HashPassword(req.Password)
	if err != nil {
		return errs.NewStack(err)
	}

	user, err := h.getDB().CreateUser(req.Login, cryptedPwd, req.Role)
	if err != nil {
		return errs.NewStack(err)
	}

	return c.JSON(http.StatusOK, transport.CreateUserResponse{User: user})
}

func (h *Handlers) UpdateUser(c echo.Context) error {
	req := &transport.UpdateUserRequest{}
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return errs.NewStack(err)
	}

	if err := transport.BindAndValidate(c, req); err != nil {
		return errs.NewStack(err)
	}

	cryptedPwd, err := crypto.HashPassword(req.Password)
	if err != nil {
		return errs.NewStack(err)
	}

	user, err := h.getDB().UpdateUser(uint(id), req.Login, cryptedPwd, req.Role)
	if err != nil {
		return errs.NewStack(err)
	}

	return c.JSON(http.StatusOK, transport.UpdateUserResponse{User: user})
}

func (h *Handlers) DeleteUser(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return errs.NewStack(err)
	}

	if err := h.getDB().DeleteUser(uint(id)); err != nil {
		return errs.NewStack(err)
	}
	if err := h.getDB().DeleteSessionsForUser(uint(id)); err != nil {
		return errs.NewStack(err)
	}

	return c.JSON(http.StatusOK, transport.BaseResponse{})
}
