package userstorage

import (
	"context"
	"mine-chat/common"
	usermodel "mine-chat/module/user/model"
)

func (s *sqlStore) Create(context context.Context, data *usermodel.UserCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
