package chatmessagebiz

import (
	"context"
	modelchatmessage "mine-chat/module/chatmessage/model"
)

type postChatMessageStore interface {
	Create(ctx context.Context, data *modelchatmessage.ChatMessage) error
}

type postChatMessageBiz struct {
	store postChatMessageStore
}

func NewPostChatMessageBiz(store postChatMessageStore) *postChatMessageBiz {
	return &postChatMessageBiz{store: store}
}

func (biz postChatMessageBiz) PostNewChatMessageBiz(context context.Context, data *modelchatmessage.ChatMessage) error {
	if err := biz.store.Create(context, data); err != nil {
		return err
	}

	return nil
}
