package handlers

import (
	"github.com/gophers0/users/internal/model"
	"github.com/gophers0/users/internal/transport"
	"github.com/gophers0/users/pkg/errs"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handlers) SearchUser(c echo.Context) error {
	req := &transport.SearchRequest{}
	err := transport.BindAndValidate(c, req)
	if err != nil {
		return errs.NewStack(errs.RequestValidationError.AddInfoErrMessage(err))
	}

	user := c.Get(transport.CtxUserKey).(*model.User)
	records, count, err := h.getDB().SearchUsers(req.Login, user.Role == model.UserRoleAdmin, req.Limit, req.Offset)
	if err != nil {
		return errs.NewStack(err)
	}

	return c.JSON(http.StatusOK, transport.SearchResponse{
		Count:   count,
		Records: records,
	})
}
