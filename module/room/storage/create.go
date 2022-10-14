package roomstorage

import (
	"context"
	"mine-chat/common"
	roommodel "mine-chat/module/room/model"
)

func (s *sqlStore) Create(context context.Context, data *roommodel.ChatRoom) error {
	if err := s.db.Table(roommodel.ChatRoom{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
