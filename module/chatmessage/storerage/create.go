package chatmessagestorerage

import (
	"context"
	"mine-chat/common"
	modelchatmessage "mine-chat/module/chatmessage/model"
)

func (s sqlStore) Create(ctx context.Context, data *modelchatmessage.ChatMessage) error {
	if err := s.db.Table(modelchatmessage.ChatMessage{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
