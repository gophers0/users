package postgres

import (
	"errors"
	"github.com/gophers0/users/internal/model"
	"github.com/gophers0/users/pkg/errs"
	"github.com/jinzhu/gorm"
)

func (r *Repo) FindUserByLogin(login string) (*model.User, error) {
	user := &model.User{}
	if err := r.DB.Where("login = ?", login).First(user).Error; err != nil {
		return nil, errs.NewStack(err)
	}

	return user, nil
}

func (r *Repo) CreateUser(login, password, role string) (*model.User, error) {
	_, err := r.FindUserByLogin(login)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// create user
		user := &model.User{
			Model:    model.Model{},
			Login:    login,
			Password: password,
			Role:     role,
		}
		if err := r.DB.Create(user).Error; err != nil {
			return nil, errs.NewStack(err)
		}
	}
	return nil, errs.NewStack(errs.UserAlreadyExists)
}

func (r *Repo) UpdateUser(userId uint, login, password, role string) (*model.User, error) {
	user := &model.User{
		Login:    login,
		Password: password,
		Role:     role,
	}
	count := 0
	if err := r.DB.Model(user).Where(userId).Count(&count).Error; err != nil {
		return nil, errs.NewStack(err)
	}
	if err := r.DB.Where(userId).UpdateColumns(user).Error; err != nil {
		return nil, errs.NewStack(err)
	}

	return user, nil
}

func (r *Repo) DeleteUser(userId uint) error {
	return errs.NewStack(r.DB.Where("id = ?", userId).Delete(&model.User{}).Error)
}
