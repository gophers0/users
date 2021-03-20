package postgres

import (
	"github.com/gophers0/users/internal/model"
	"github.com/gophers0/users/pkg/errs"
	"github.com/gophers0/users/pkg/hexdigest"
)

func (r *Repo) GetSessionForUser(userId uint) (*model.Session, error) {
	user := &model.User{}
	if err := r.DB.First(user, userId).Error; err != nil {
		return nil, errs.NewStack(err)
	}

	session := &model.Session{
		User: *user,
	}
	newToken, err := hexdigest.HexDigest()
	if err != nil {
		return nil, errs.NewStack(err)
	}

	if err := r.DB.Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Preload("User").
		Where("user_id = ?", userId).
		Attrs(model.Session{Token: newToken}).
		FirstOrCreate(session).Error; err != nil {
		return nil, errs.NewStack(err)
	}

	return session, nil
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
