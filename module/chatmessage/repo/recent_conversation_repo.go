package chatmessagerepo

import (
	"context"
	"mine-chat/common"
	modelchatmessage "mine-chat/module/chatmessage/model"
	roommodel "mine-chat/module/room/model"
)

type ChatMessageStore interface {
	ListChatMessageRecentConversation(ctx context.Context,
		roomIds []int,
		filter *modelchatmessage.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]modelchatmessage.ChatMessage, error)
}

type ChatRoomStore interface {
	FindRoomWithUserId(ctx context.Context, userId int) ([]roommodel.ChatRoom, error)
}

type RecentConversationRepo struct {
	chatRoomStore    ChatRoomStore
	chatMessageStore ChatMessageStore
}

func NewRecentConversationRepo(chatRoomStore ChatRoomStore, chatMessageStore ChatMessageStore) *RecentConversationRepo {
	return &RecentConversationRepo{
		chatRoomStore:    chatRoomStore,
		chatMessageStore: chatMessageStore,
	}
}

func (repo *RecentConversationRepo) ListRecentConversation(ctx context.Context,
	userId int,
	filter *modelchatmessage.Filter,
	paging *common.Paging,
) ([]modelchatmessage.ChatMessage, error) {
	rooms, err := repo.chatRoomStore.FindRoomWithUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	roomIds := make([]int, len(rooms))
	for i := range rooms {
		roomIds[i] = rooms[i].Id
	}

	data, err := repo.chatMessageStore.ListChatMessageRecentConversation(ctx, roomIds, filter, paging, "PostByUser")
	if err != nil {
		return nil, err
	}

	return data, nil
}
