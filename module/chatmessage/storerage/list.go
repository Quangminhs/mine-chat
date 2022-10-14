package chatmessagestorerage

import (
	"context"
	"mine-chat/common"
	modelchatmessage "mine-chat/module/chatmessage/model"
)

func (s *sqlStore) ListChatMessageByCondition(ctx context.Context, conditions map[string]interface{},
	filter *modelchatmessage.Filter,
	paging *common.Paging, moreKeys ...string,
) ([]modelchatmessage.ChatMessage, error) {
	var result []modelchatmessage.ChatMessage
	db := s.db

	db = db.Table(modelchatmessage.ChatMessage{}.TableName()).Where(conditions).Where("status in (1)")

	if v := filter; v != nil {
		if v.PostByUser > 0 {
			db = db.Where("post_by_user = ?", v.PostByUser)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil

}
