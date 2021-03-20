package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gophers0/users/internal/transport"
	"github.com/gophers0/users/pkg/errs"
	"github.com/labstack/echo"
)

func (mw *Middleware) Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			// Get Authorization header
			userId, token, err := parseAuthorizationHeader(c.Request())
			if err != nil {
				return errs.NewStack(err)
			}
			session, err := mw.repo.GetSessionForUser(uint(userId))
			if err != nil {
				return errs.InvalidToken
			}

			if session.Token != token {
				return errs.InvalidToken
			}

			c.Set(transport.CtxSessionKey, session)
			c.Set(transport.CtxUserKey, &session.User)
			return next(c)
		}
	}
}

func parseAuthorizationHeader(req *http.Request) (int, string, error) {
	header := req.Header.Get(transport.AuthorizationHeader)
	if len(header) == 0 {
		return 0, "", errs.AuthorizationHeaderMissing
	}

	segs := strings.Split(header, ":")

	if len(segs) != 2 {
		return 0, "", errs.InvalidAuthorizationHeader
	}

	userId, err := strconv.Atoi(segs[0])
	if err != nil {
		return 0, "", errs.InvalidAuthorizationHeader
	}

	if len(segs[1]) == 0 {
		return 0, "", errs.InvalidAuthorizationHeader
	}

	return userId, segs[1], nil
}
