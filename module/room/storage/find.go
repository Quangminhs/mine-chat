package roomstorage

import (
	"context"
	"mine-chat/common"
	roommodel "mine-chat/module/room/model"
)

func (s *sqlStore) FindRoomWithUserIds(ctx context.Context, userIds string) (*roommodel.ChatRoomCreate, error) {
	var data roommodel.ChatRoomCreate
	err := s.db.Table(roommodel.ChatRoom{}.TableName()).Where("user_ids = ?", userIds).First(&data).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
