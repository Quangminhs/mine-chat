package userstorage

import (
	"context"
	usermodel "mine-chat/module/user/model"
)

func (s *sqlStore) ListUser(context context.Context) ([]usermodel.UserSimple, error) {
	var result []usermodel.UserSimple
	if err := s.db.Table(usermodel.User{}.TableName()).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
