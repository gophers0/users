package postgres

import (
	"github.com/gophers0/users/internal/model"
	"github.com/gophers0/users/pkg/errs"
)

func (r *Repo) GetSessionTokenForUser(userId uint) (string, error) {
	user := &model.User{}
	if err := r.DB.First(user, userId).Error; err != nil {
		return "", errs.NewStack(err)
	}

	session := &model.Session{}
	err := r.DB.FirstOrCreate(session, "user_id = ?", userId).Error
	if err != nil {
		return "", errs.NewStack(err)
	}

	return session.Token, nil
}

func (r *Repo) CheckSessionTokenForUser(token string, userId uint) (*model.Session, error) {
	session := &model.Session{}
	err := r.DB.Preload("User").Where("token = ? AND user_id = ?", token, userId).First(session).Error
	if err != nil {
		return nil, errs.NewStack(err)
	}
	return session, err
}

func (r *Repo) DeleteSessionsForUser(userId uint) error {
	return r.DB.Where("user_id = ?", userId).Delete(&model.Session{}).Error
}
