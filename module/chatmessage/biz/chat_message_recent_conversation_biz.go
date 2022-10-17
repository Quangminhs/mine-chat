package chatmessagebiz

import (
	"context"
	"mine-chat/common"
	modelchatmessage "mine-chat/module/chatmessage/model"
)

type RecentConversationRepo interface {
	ListRecentConversation(ctx context.Context,
		userId int,
		filter *modelchatmessage.Filter,
		paging *common.Paging,
	) ([]modelchatmessage.ChatMessage, error)
}

type chatMessageRecentConversationBiz struct {
	repo RecentConversationRepo
}

func NewChatMessageRecentConversationBiz(repo RecentConversationRepo) *chatMessageRecentConversationBiz {
	return &chatMessageRecentConversationBiz{
		repo: repo,
	}
}

func (biz *chatMessageRecentConversationBiz) GetListMessageRecentConversation(ctx context.Context, userId int,
	filter *modelchatmessage.Filter,
	paging *common.Paging) ([]modelchatmessage.ChatMessage, error) {
	data, err := biz.repo.ListRecentConversation(ctx, userId, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(modelchatmessage.EntityName, err)
	}

	return data, nil
}
