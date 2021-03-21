package postgres

import (
	"github.com/gophers0/users/internal/model"
	"github.com/gophers0/users/pkg/errs"
)

func (r *Repo) SearchUsers(login string, isAdmin bool, limit, offset int) ([]*model.User, int, error) {
	whereString := " lower (users.login) like ?"
	whereParams := []interface{}{login}

	if isAdmin {
		whereParams = []interface{}{login + "%"}
	}

	db := r.DB.Model(model.User{}).
		Where(whereString, whereParams...)

	var count int64
	if err := db.Debug().Count(&count).Error; err != nil {
		return nil, 0, errs.NewStack(err)
	}

	if limit == 0 || limit > 99 {
		limit = 99
	}
	if offset < 0 {
		offset = 0
	}

	var res []*model.User
	err := db.
		Limit(limit).
		Offset(offset).
		Find(&res).
		Error

	return res, int(count), errs.NewStack(err)
}
