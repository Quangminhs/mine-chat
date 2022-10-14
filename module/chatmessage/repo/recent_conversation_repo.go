package chatmessagerepo

import (
	"context"
	"mine-chat/common"
	modelchatmessage "mine-chat/module/chatmessage/model"
)

type ChatMessageStore interface {
	ListChatMessageByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *modelchatmessage.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]modelchatmessage.ChatMessage, error)
}

type ChatRoomStore interface {
	FindRoomWithUserIds(ctx context.Context, userIds string) (int, error)
}

type recentConversationRepo struct {
	chatRoomStore    ChatRoomStore
	chatMessageStore ChatMessageStore
}

func NewRecentConversationRepo(chatRoomStore ChatRoomStore, chatMessageStore ChatMessageStore) *recentConversationRepo {
	return &recentConversationRepo{
		chatRoomStore:    chatRoomStore,
		chatMessageStore: chatMessageStore,
	}
}

func (repo *recentConversationRepo) ListRecentConversation(ctx context.Context,
	userId int,
	filter *modelchatmessage.Filter,
	paging *common.Paging,
) ([]modelchatmessage.ChatMessage, error) {
	//epo.chatRoomStore.FindRoomWithUserIds(ctx, userId)
	return nil, nil
}
