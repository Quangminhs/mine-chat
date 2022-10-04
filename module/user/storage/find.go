package userstorage

import (
	"context"
	"mine-chat/common"
	usermodel "mine-chat/module/user/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(context context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	var data usermodel.User
	err := s.db.Where(condition).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}

func (s *sqlStore) FindUser(context context.Context, condition map[string]interface{}, moreKeys ...string) (*usermodel.User, error) {
	return s.FindDataWithCondition(context, condition, moreKeys...)
}
