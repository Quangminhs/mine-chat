package chatmessagebiz

import (
	"context"
	"mine-chat/common"
	modelchatmessage "mine-chat/module/chatmessage/model"
)

type ChatMessageListStore interface {
	ListChatMessageByCondition(ctx context.Context, conditions map[string]interface{},
		filter *modelchatmessage.Filter,
		paging *common.Paging, moreKeys ...string,
	) ([]modelchatmessage.ChatMessage, error)
}

type chatMessageListBiz struct {
	store ChatMessageListStore
}

func NewChatMessageListBiz(store ChatMessageListStore) *chatMessageListBiz {
	return &chatMessageListBiz{
		store: store,
	}
}

func (biz *chatMessageListBiz) ListChatMessageList(ctx context.Context,
	conditions map[string]interface{},
	filter *modelchatmessage.Filter,
	paging *common.Paging) ([]modelchatmessage.ChatMessage, error) {
	result, err := biz.store.ListChatMessageByCondition(ctx, conditions, filter, paging, "PostByUser")

	if err != nil {
		return nil, common.ErrCannotListEntity(modelchatmessage.EntityName, err)
	}

	return result, nil
}
