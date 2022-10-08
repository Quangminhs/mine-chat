package roombiz

import (
	"context"
	roommodel "mine-chat/module/room/model"
	"sort"
	"strconv"
	"strings"
)

type chatRoomStore interface {
	Create(context context.Context, data *roommodel.ChatRoomCreate) error
	FindRoomWithUserIds(ctx context.Context, userIds string) (*roommodel.ChatRoomCreate, error)
}

type chatRoomBiz struct {
	store chatRoomStore
}

func NewChatRoom(store chatRoomStore) *chatRoomBiz {
	return &chatRoomBiz{
		store: store,
	}
}

func (biz chatRoomBiz) InitiateChat(ctx context.Context, data *roommodel.ChatRoomCreate) error {
	maps := strings.Split(data.UserIds, ",")
	sort.Slice(maps, func(i, j int) bool {
		x, er := strconv.Atoi(maps[i])
		if er != nil {
			x = 0
		}
		y, er := strconv.Atoi(maps[i])
		if er != nil {
			y = 0
		}

		return x < y
	})

	userIds := strings.Join(maps[:], ",")
	data.UserIds = userIds

	chatRoom, err := biz.store.FindRoomWithUserIds(ctx, data.UserIds)

	//Nếu không tìm thấy thì tạo mới room chat
	if err != nil {
		err2 := biz.store.Create(ctx, data)
		if err2 != nil {
			return err2
		}

		return nil
	}

	//Trả về id Room cho client
	data.Id = chatRoom.Id
	return nil
}
