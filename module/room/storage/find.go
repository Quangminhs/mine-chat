package roomstorage

import (
	"context"
	"mine-chat/common"
	roommodel "mine-chat/module/room/model"
)

func (s *sqlStore) FindRoomWithUserIds(ctx context.Context, userIds []int) (int, error) {
	var data roommodel.ChatRoom
	err := s.db.Table(roommodel.ChatRoom{}.TableName()).Where("user_ids = JSON_ARRAY ?", userIds).First(&data).Error
	if err != nil {
		return -1, common.ErrDB(err)
	}

	return data.Id, nil
}

func (s *sqlStore) FindRoomWithUserId(ctx context.Context, userId int) ([]roommodel.ChatRoom, error) {
	var data []roommodel.ChatRoom
	err := s.db.
		Raw(" SELECT * FROM chat_room WHERE JSON_CONTAINS(user_ids, '1','$');").Scan(&data).Error
	if err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
