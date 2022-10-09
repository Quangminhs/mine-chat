package roomstorage

import (
	"context"
	"mine-chat/common"
	roommodel "mine-chat/module/room/model"
)

func (s *sqlStore) FindRoomWithUserIds(ctx context.Context, userIds string) (int, error) {
	var data roommodel.ChatRoom
	err := s.db.Table(roommodel.ChatRoom{}.TableName()).Where("user_ids = ?", userIds).First(&data).Error
	if err != nil {
		return -1, common.ErrDB(err)
	}

	return data.Id, nil
}
