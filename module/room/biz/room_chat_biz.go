package roombiz

import (
	"context"
	"mine-chat/common"
	roommodel "mine-chat/module/room/model"
	"sort"
	"strconv"
	"strings"
)

type chatRoomStore interface {
	Create(context context.Context, data *roommodel.ChatRoom) error
	FindRoomWithUserIds(ctx context.Context, userIds string) (int, error)
}

type chatRoomBiz struct {
	store chatRoomStore
}

func NewChatRoom(store chatRoomStore) *chatRoomBiz {
	return &chatRoomBiz{
		store: store,
	}
}

func (biz chatRoomBiz) InitiateChat(ctx context.Context, data *roommodel.ChatRoomRequest) error {
	maps := strings.Split(data.UserIds, ",")
	for i := range maps {
		uid, err := common.FromBase58(maps[i])
		if err != nil {
			return nil
		}
		maps[i] = strconv.Itoa(int(uid.GetLocalID()))
	}
	sort.Strings(maps)
	//sort.Slice(maps, func(i, j int) bool {
	//	x, er := strconv.Atoi(maps[i])
	//	if er != nil {
	//		x = 0
	//	}
	//	y, er := strconv.Atoi(maps[i])
	//	if er != nil {
	//		y = 0
	//	}
	//
	//	return x < y
	//})

	userIds := strings.Join(maps[:], ",")
	//data.UserIds = userIds

	chatRoomId, err := biz.store.FindRoomWithUserIds(ctx, userIds)

	//Nếu không tìm thấy thì tạo mới room chat
	if err != nil {
		uid, err2 := common.FromBase58(data.ChatInitiator)
		if err2 != nil {
			return nil
		}
		chatInitiator := strconv.Itoa(int(uid.GetLocalID()))
		dataCreate := roommodel.ChatRoom{
			UserIds:       userIds,
			ChatInitiator: chatInitiator,
			Type:          "public",
		}
		err3 := biz.store.Create(ctx, &dataCreate)
		if err3 != nil {
			return err2
		}
		data.Id = common.GenUID(dataCreate.Id, common.DB_TYPE_ROOM)
		return nil
	}

	//Trả về id Room cho client
	data.Id = common.GenUID(chatRoomId, common.DB_TYPE_ROOM)
	return nil
}
